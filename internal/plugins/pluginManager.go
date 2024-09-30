package plugins

import (
    "fmt"
    "sync"
    "gowiki/internal/types"
    "github.com/labstack/echo/v4"
)

type Manager struct {
    plugins map[string]types.Plugin
    mu      sync.RWMutex
}

func NewManager() *Manager {
    return &Manager{
        plugins: make(map[string]types.Plugin),
    }
}

func (m *Manager) RegisterPlugin(p types.Plugin) error {
    m.mu.Lock()
    defer m.mu.Unlock()

    if _, exists := m.plugins[p.Name()]; exists {
        return fmt.Errorf("plugin %s already registered", p.Name())
    }

    m.plugins[p.Name()] = p
    return p.Initialize()
}

func (m *Manager) GetPlugin(name string) (types.Plugin, bool) {
    m.mu.RLock()
    defer m.mu.RUnlock()

    p, ok := m.plugins[name]
    return p, ok
}

func (m *Manager) AllPlugins() []types.Plugin {
    m.mu.RLock()
    defer m.mu.RUnlock()

    plugins := make([]types.Plugin, 0, len(m.plugins))
    for _, p := range m.plugins {
        plugins = append(plugins, p)
    }
    return plugins
}

func (m *Manager) ShutdownAll() error {
    m.mu.Lock()
    defer m.mu.Unlock()

    for _, p := range m.plugins {
        if err := p.Shutdown(); err != nil {
            return fmt.Errorf("error shutting down plugin %s: %w", p.Name(), err)
        }
    }
    return nil
}

func (m *Manager) RegisterHandlers(e *echo.Echo) {
    for _, p := range m.AllPlugins() {
        for route, handler := range p.Handlers() {
            e.GET(route, handler)
        }
    }
}

func (m *Manager) GetTemplateData() map[string]interface{} {
    data := make(map[string]interface{})
    for _, p := range m.AllPlugins() {
        for k, v := range p.TemplateData() {
            data[k] = v
        }
    }
    return data
}
package CardCreator

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "path/filepath"
)

type CardLayout struct {
    Elements []CardElement `json:"elements"`
}

type CardElement struct {
    Type     string `json:"type"`
    Position struct {
        X int `json:"x"`
        Y int `json:"y"`
    } `json:"position"`
}

type DocumentData struct {
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Type        string   `json:"type"`
    Category    string   `json:"category"`
    Status      string   `json:"status"`
    Tags        []string `json:"tags"`
    Content     string   `json:"content"`
    Project     string   `json:"project"`
    Path        string   `json:"path"`
}

func SaveCardLayout(layout CardLayout) error {
    data, err := json.Marshal(layout)
    if err != nil {
        return err
    }
    return ioutil.WriteFile("data/.settings/card_layout.json", data, 0644)
}

func LoadCardLayout() (CardLayout, error) {
    var layout CardLayout
    data, err := ioutil.ReadFile("data/.settings/card_layout.json")
    if err != nil {
        if os.IsNotExist(err) {
            return CardLayout{}, nil
        }
        return layout, err
    }
    err = json.Unmarshal(data, &layout)
    return layout, err
}

func LoadDocumentData(path string) (DocumentData, error) {
    var doc DocumentData
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return doc, err
    }
    err = json.Unmarshal(data, &doc)
    if err != nil {
        return doc, err
    }
    doc.Project = filepath.Base(filepath.Dir(filepath.Dir(path)))
    doc.Path = path
    return doc, nil
}
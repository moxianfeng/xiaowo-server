package model

import (
)

var allModels []interface{};

func GetModels() []interface{} {
    return allModels;
}

func RegisterObject(v interface{}) {
    if nil == allModels {
        allModels = make([]interface{}, 0);
    }

    allModels = append(allModels, v);
}


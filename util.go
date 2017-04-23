package main

func Contains(str string, arr []string) bool {
    for _, item := range arr {
        if str == item {
            return true
        }
    }

    return false
}
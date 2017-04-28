package main

/*
Check if string array contains string
 */
func Contains(str string, arr []string) bool {
    for _, item := range arr {
        if str == item {
            return true
        }
    }

    return false
}

/*
Get keys of a map as array
 */
func Keys(dict map[string]string) (keys []string) {
    for k := range dict {
        keys = append(keys, k)
    }

    return keys
}

/*
Get values of a map as array
 */
func Values(dict map[string]string) (values []string) {
    for _,v := range dict {
        values = append(values, v)
    }

    return values
}
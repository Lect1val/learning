package main

func longestCommonPrefix(strs []string) string {
    //edge case empty input
    if len(strs) == 0 {
        return ""
    }

    //first assume that the first string is the longest common prefix
    prefix := strs[0]

    for i := 1; i < len(strs); i++ {

        // while the current strs[i] doesn't start with the prefix
        for len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix {
            //shorten the prefix by one charactor
            prefix = prefix[:len(prefix)-1]

            // if prefix become empty, then there is no common prefix among the input
            if prefix == "" {
                return ""
            }
        }
    }

    return prefix
}

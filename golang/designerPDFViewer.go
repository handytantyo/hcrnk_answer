package golang

import "strings"

func designerPdfViewer(h []int32, word string) int32 {
    // make it all words lower case
    word = strings.ToLower(word)
    wordLength := int32(len(word))

    // find max height
    var maxHeight int32
    for _, wordInput := range word {
        height := int32(h[int(wordInput)-97])
        if maxHeight < height {
            maxHeight = height
		}
    }

    // calculate the result
    result := wordLength * maxHeight
    return result
}
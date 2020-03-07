package solutions

import "fmt"

type WordLevel struct {
	word  string
	level int
}

func isWordInWordList(word string, wordList *[]string) bool {
	for _, wl := range *wordList {
		if wl == word {
			return true
		}
	}

	return false
}

func generatePatterns(word string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := range word {
			wordRunes := []rune(word)
			wordRunes[i] = '*'
			channel <- string(wordRunes)
		}
		close(channel)
	}()
	return channel
}

func LadderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) > 0 && isWordInWordList(endWord, &wordList) {
		patterns := make(map[string][]string)
		for _, word := range wordList {
			for pattern := range generatePatterns(word) {
				patterns[pattern] = append(patterns[pattern], word)
			}
		}

		queue := []WordLevel{{beginWord, 1}}
		visited := make(map[string]bool)

		for len(queue) > 0 {
			currentWord, currentLevel := queue[0].word, queue[0].level
			queue = queue[1:]

			for cwPattern := range generatePatterns(currentWord) {
				if patterns[cwPattern] == nil {
					continue
				}

				for _, wordRelative := range patterns[cwPattern] {
					if wordRelative == endWord {
						return currentLevel + 1
					} else if !visited[wordRelative] {
						visited[wordRelative] = true
						queue = append(queue, WordLevel{
							wordRelative,
							currentLevel + 1,
						})
					}
				}
			}

			fmt.Println(currentWord)
		}
	}

	return 0
}

package promptfuse

import (
	"os"
	"strings"

	"github.com/pkoukk/tiktoken-go"
)

const DefaultDictEntryCost = 4

type TokenInfo struct {
	Index int
	ID    int
	Text  string
}

type Config struct {
	FilePath     string
	ModelName    string
	PhraseLength int
}

// Analyze tokenizes the input file, identifies repeated phrases of the configured length,
// and computes the potential token savings if those phrases were replaced using a dictionary encoding.
//
// It returns:
// - A map of phrases to their individual net savings (after subtracting the dictionary entry cost)
// - The total number of tokens saved across all repeated phrases
// - An error if file reading or tokenization fails
func (cfg *Config) Analyze() (map[string]int, int, error) {
	tokens, err := Tokenize(cfg.FilePath, cfg.ModelName)
	if err != nil {
		return nil, 0, err
	}

	repeats := CountPhraseRepetition(tokens, cfg.PhraseLength)
	savings := ComputeSavingsByPhrase(repeats, DefaultDictEntryCost)
	total := ComputeTotalSavings(repeats, DefaultDictEntryCost)

	return savings, total, nil
}

// Tokenize grabs any text file, uses the tokenizer,
// and returns the tokens in structured format.
func Tokenize(filePath, modelName string) ([]TokenInfo, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	enc, err := tiktoken.EncodingForModel(modelName)
	if err != nil {
		return nil, err
	}

	text := string(data)
	tokens := enc.Encode(text, nil, nil)
	tokenInfos := make([]TokenInfo, len(tokens))

	for i, token := range tokens {
		tokenStr := enc.Decode([]int{token})
		tokenInfos[i] = TokenInfo{
			Index: i,
			ID:    token,
			Text:  tokenStr,
		}
	}

	return tokenInfos, nil
}

// BuildPhrase constructs a phrase of `length` tokens, starting at position `i`
func BuildPhrase(tokens []TokenInfo, i int, length int) (string, bool) {
	if i+length > len(tokens) {
		return "", false
	}

	var parts []string
	for j := range length {
		parts = append(parts, tokens[i+j].Text)
	}

	var phrase strings.Builder
	for j, part := range parts {
		if j > 0 {
			prev := parts[j-1]
			curr := part

			if isWord(curr) &&
				!strings.HasSuffix(prev, " ") &&
				!strings.HasPrefix(curr, " ") &&
				!startsWithSymbol(curr) {
				phrase.WriteString(" ")
			}
		}
		phrase.WriteString(part)
	}

	return phrase.String(), true
}

// CountPhraseRepetition counts the frequency of all phrases of the given token length.
func CountPhraseRepetition(tokenInfos []TokenInfo, phraseLength int) map[string]int {
	repetitionMap := make(map[string]int)

	for i := 0; i <= len(tokenInfos)-phraseLength; i++ {
		phrase, ok := BuildPhrase(tokenInfos, i, phraseLength)
		if !ok {
			continue
		}
		repetitionMap[phrase]++
	}

	return repetitionMap
}

// ComputeSavingsByPhrase returns a map of phrases and their net token savings
func ComputeSavingsByPhrase(countMap map[string]int, dictEntryCost int) map[string]int {
	savingsMap := make(map[string]int)
	for phrase, count := range countMap {
		netSaved := count - dictEntryCost
		if netSaved > 0 {
			savingsMap[phrase] = netSaved
		}
	}
	return savingsMap
}

// ComputeTotalSavings calculates total savings across all repeated phrases
func ComputeTotalSavings(countMap map[string]int, dictEntryCost int) int {
	total := 0
	for _, count := range countMap {
		saved := count - dictEntryCost
		if saved > 0 {
			total += saved
		}
	}
	return total
}

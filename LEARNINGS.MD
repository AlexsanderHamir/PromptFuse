### 🔍 Tokenization Insights

1. **Capitalization can affect tokenization**  
   If your prompt uses **capitalized words that aren't commonly capitalized**, they may be broken into multiple tokens. For example, `"Designer"` might be tokenized differently than `"designer"`.

2. **Uncommon words may split into multiple tokens**  
   Words that are **rare or out-of-vocabulary** may be split into smaller, more frequent subwords. For instance, `"visioneering"` might become `"vision"` + `"eering"`.

3. **Whole words count as single tokens—if not split**  
   As long as a word is **recognized in full** by the tokenizer, it will typically count as one token.

4. **Long or rare words may be broken into popular substrings**  
   Tokenizers are designed to favor **frequent substrings**. If a word is long or unusual, it's more efficient (and common) for it to be tokenized as multiple **more common parts** rather than as one unfamiliar token.

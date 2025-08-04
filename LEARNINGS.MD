### 🔍 Tokenization Insights

1. **Capitalization affects tokenization**  
   Words with uncommon capitalization (e.g., `"Designer"` vs. `"designer"`) may be split into multiple tokens.

2. **Rare or uncommon words split into subwords**  
   Unfamiliar words (like `"visioneering"`) are often broken into smaller frequent parts (e.g., `"vision"` + `"eering"`).

3. **Common whole words usually form single tokens**  
   If a word is recognized fully by the tokenizer, it counts as one token.

4. **Long or rare words break into frequent substrings**  
   Tokenizers favor frequent substrings over unfamiliar whole tokens, so long or unusual words get split accordingly.


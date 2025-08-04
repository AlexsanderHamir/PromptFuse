### 🔍 Tokenization Insights

1. **Capitalization affects tokenization**  
   Words with uncommon capitalization (e.g., `"Designer"` vs. `"designer"`) may be split into multiple tokens.

2. **Rare or uncommon words split into subwords**  
   Unfamiliar words (like `"visioneering"`) are often broken into smaller frequent parts (e.g., `"vision"` + `"eering"`).

3. **Common whole words usually form single tokens**  
   If a word is recognized fully by the tokenizer, it counts as one token.

4. **Long or rare words break into frequent substrings**  
   Tokenizers favor frequent substrings over unfamiliar whole tokens, so long or unusual words get split accordingly.

5. **Shared Codes for Compression**
   Sending the full dictionary encoding map with every prompt **adds overhead** and can actually **increase token usage**. However, if you compute the dictionary once and **reuse it across multiple queries** — by embedding it in the system prompt or agent memory — you can **significantly reduce token cost**, especially in domains with repetitive phrasing or specialized terminology. This approach works best when interacting with models on focused tasks, where the same phrases and structures occur frequently.

6. **Number Tokenization**
   Numbers are often split into multiple tokens rather than remaining a single token, which can increase token usage.

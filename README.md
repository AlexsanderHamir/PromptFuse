# TokenSpan: A Simple Token Compression Technique

## 🔢 THE FORMULA: `Saved = (2 × X) - (X + 4)`

TokenSpan is based on a straightforward idea:
**compress token usage by substituting frequently repeated phrases with a single-token code**.

For example, consider a common two-token phrase like `"Microsoft Designer"`.
If it appears **X** times in a prompt, the original token cost is:

```
Original cost: 2 × X tokens
```

By assigning a single-token alias (e.g., `"§0"`), each occurrence now costs **1 token**, plus a small dictionary overhead to define this mapping:

```
New cost: X (usages) + 4 (dictionary entry overhead)
```

So, the total token savings are:

```
Saved = (2 × X) - (X + 4)
```

For instance, if `"Microsoft Designer"` appears 15 times:

```
Saved = (2 × 15) - (15 + 4) = 30 - 19 = 11 tokens saved
```

Over many repeated phrases, this can add up significantly. You might think such repetition is rare, but common short phrases like `", and"`—which counts as two tokens—occur frequently in large prompts, contributing to meaningful savings.

### 🎯 Why focus on two-token phrases?

- Single-token replacements don’t reduce token counts.
- Phrases longer than two tokens save more but appear less frequently.
- Two-token phrases strike a balance — common enough to compress, large enough to matter.

### 📦 Dictionary Encoding Overhead

Each dictionary entry requires a small fixed token cost:

- **1 token** for the replacement code (e.g., `§a`, `@b`)
  _Note: Using digits can increase token splits, raising overhead._

- **1 token** for the separator (e.g., `→`, `:`)

- **2 tokens** for the original phrase

**Total overhead per dictionary entry: approximately 4 tokens**

You start to save tokens only when a phrase appears more than 4 times, and every additional repetition adds to your savings.

## 🚀 Example Compression Potential

- Raw prompt size: 8,019 tokens
- Target size (optimized): 7,138 tokens
- Savings: 881 tokens (\~11.0% reduction)

This approach is a concept for efficient prompt token compression and could be extended or integrated into tooling for token-aware prompt optimization.

## 🛠️ Small API to Estimate Compression Benefits

TokenSpan offers a simple API to analyze any text file and estimate whether dictionary compression would save tokens. It identifies frequently repeated phrases and computes potential token savings, helping you decide if applying dictionary encoding is worthwhile before transforming your prompts.

## 🔁 Shared Codes for Compression

Sending the full dictionary encoding map with every prompt **adds overhead** and can actually **increase token usage**. However, if you compute the dictionary once and **reuse it across multiple queries** — by embedding it in the system prompt or agent memory — you can **significantly reduce token cost**, especially in domains with repetitive phrasing or specialized terminology.

This approach works best when interacting with models on focused tasks, where the same phrases and structures occur frequently.

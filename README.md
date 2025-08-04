## 🔢 THE FORMULA: `Saved = (2 × X) - (X + 4)`

**PromptFuse** is built on a simple but powerful idea: compressing token usage by substituting **frequently repeated phrases** with a single token code.

Take a common two-word phrase like `"Microsoft Designer"`, which typically spans **2 tokens**. If it appears **X** times in a prompt, it costs:

```
Original cost: 2 × X tokens
```

By assigning it a one-token alias (e.g. `"§0"`), each occurrence now costs **1 token**, plus a small **dictionary overhead** to define the mapping:

```
New cost: X (usages) + 4 (dictionary entry)
```

So the total **token savings** becomes:

```
Saved = (2 × X) - (X + 4)
```

For example, if `"Microsoft Designer"` appears 15 times:

```
Saved = (2 × 15) - (15 + 4) = 30 - 19 = ✅ 11 tokens saved
```

This can add up quickly in long prompts.

### 🎯 Why Focus on Two-Token Phrases?

- Replacing **1 token** with another doesn’t help.
- Replacing **3+ tokens** offers more savings but occurs less often.
- **2-token sequences** hit the sweet spot: they're common enough to find and compress, and large enough to matter.

PromptFuse targets these high-frequency 2-token spans for maximum impact with minimal complexity.

### 📦 Dictionary Encoding Overhead

Each dictionary entry incurs a small token cost, composed of:

- **1 token** for the **replacement code** (e.g., `§a`, `@b`, etc.)
  _Note:_ Using numeric characters in codes may cause tokenization to split tokens more frequently, potentially increasing dictionary cost.

- **1 token** for the **separator** symbol (such as `→` or `:`)

- **2 tokens** for the **original two-word phrase** being replaced

**Total overhead per dictionary entry:** approximately **4 tokens**

**Total: \~4 tokens** per dictionary entry

You start saving tokens once the phrase appears **more than 4 times**. After that, each extra repetition adds to your net gain.

## 🚀 Performance Goals

- **Raw Prompt Size:** 8,019 tokens
- **Target (Optimized):** 7,138 tokens
- **Token Savings:** 881 tokens (~11.0% reduction)

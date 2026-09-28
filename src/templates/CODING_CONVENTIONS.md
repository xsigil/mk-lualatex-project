# LaTeX Macro Coding Standard: Namespace & Prefix Conventions

This standard defines a macro naming scheme designed for large-scale academic books and mathematical monographs (Book class). It prevents identifier collisions in the global TeX scope while maximizing discoverability and efficiency via Vim completion (`<C-n>` / `<C-p>`), `ctags`, and `ripgrep` (`rg`).

---

## 1. Prefix Hierarchy

| Prefix | Role | Semantics & Purpose | Examples |
| :--- | :--- | :--- | :--- |
| `\mth...` | **Math Glyph / Operator** | Models, operators, symbols | `\mthEgo`, `\mthWisdom` |
| `\txt...` | **Standard Concept Text** | Standard concept names in prose | `\txtWholeSelf` |
| `\ja...`  | **Japanese Terminology**  | Japanese canonical terminology | `\jaWholeSelf` |
| `\abbr...`| **Abbreviation / Acronym**| Acronyms and shorthand tokens | `\abbrDAG` |
| `\full...`| **Expanded / Formal Text**| Fully expanded formal names | `\fullWholeSelf` |

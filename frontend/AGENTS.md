<!-- BEGIN:nextjs-agent-rules -->
# This is NOT the Next.js you know

This version has breaking changes — APIs, conventions, and file structure may all differ from your training data. Read the relevant guide in `node_modules/next/dist/docs/` before writing any code. Heed deprecation notices.
<!-- END:nextjs-agent-rules -->

## UI/UX Pro Max skill

This project uses the [UI/UX Pro Max](https://github.com/nextlevelbuilder/ui-ux-pro-max-skill) OpenCode skill for design intelligence. It is installed globally and ignored from git.

Install / update it with:

```bash
npm install -g ui-ux-pro-max-cli
uipro init --ai opencode --global
```

Generate or refresh the project design system from the repo root:

```bash
python3 ~/.opencode/skills/ui-ux-pro-max/scripts/search.py \
  "education platform learning management system" \
  --design-system --persist -p "Plato"
```

## Memory (agentmemory)

Use `memory_save` and `memory_recall` to persist and retrieve project context across sessions.

- **Always save** architectural decisions, discovered gotchas, and workflow patterns.
- Use `project: "plato"` for all memories in this project.
- Use `type` field: `architecture`, `bug`, `workflow`, `pattern`, `fact`.
- Add `concepts` and `files` for precise recall later.
- **Recall at session start** to check for existing context before investigating.

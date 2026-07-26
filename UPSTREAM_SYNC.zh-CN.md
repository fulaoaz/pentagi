# 跟进上游更新

汉化分支以 `vxcontrol/pentagi:main` 为上游基准，保留英文词典作为回退和差异参照。

## 本地同步

请先提交当前改动，再在 `zh-CN` 分支执行：

```bash
node scripts/sync-upstream.mjs
```

脚本会拉取 `origin/main`，以非快进方式合并，并运行界面英文基线检查。若上游新增了用户可见英文，合并会停在提交前，待文案接入 `en.ts` 和 `zh-CN.ts` 后再提交。只检查是否有更新时使用：

```bash
node scripts/sync-upstream.mjs --check
```

如果上游使用其他远程名称，可显式指定：

```bash
node scripts/sync-upstream.mjs --remote=upstream --branch=main
```

## 自动检查

`.github/workflows/upstream-sync.yml` 每 6 小时检查一次官方主线。有新提交时，它会在个人仓库创建或更新同步 PR；主 CI 会运行词典测试和界面英文基线检查。

界面英文基线记录尚未汉化的已知文案。新增或消失的条目都会使检查失败：

```bash
cd frontend
pnpm i18n:check
```

完成一批汉化并逐项确认后，更新基线：

```bash
pnpm i18n:baseline
```

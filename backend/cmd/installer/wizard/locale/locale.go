package locale

// Common status and UI strings
const (
	// Common status and UI strings
	UIStatistics       = "统计信息"
	UIStatus           = "状态："
	UIMode             = "模式："
	UINoConfigSelected = "未选择任何配置"
	UILoading          = "加载中..."
	UINotImplemented   = "功能尚未实现"
	UIUnsavedChanges   = "有未保存的更改"
	UIConfigSaved      = "配置已保存"

	// Status labels
	StatusEnabled       = "已启用"
	StatusDisabled      = "已禁用"
	StatusConfigured    = "已配置"
	StatusNotConfigured = "未配置"
	StatusEmbedded      = "内置"
	StatusExternal      = "外部"

	// Success/Warning messages
	MessageSearchEnginesNone       = "⚠ 未配置搜索引擎"
	MessageSearchEnginesConfigured = "✓ 已配置 %d 个搜索引擎"
	MessageDockerConfigured        = "✓ Docker 环境已配置"
	MessageDockerNotConfigured     = "⚠ Docker 环境未配置"
)

// Legend constants
const (
	LegendConfigured    = "✓ 已配置"
	LegendNotConfigured = "✗ 未配置"
)

// Common Navigation Actions (always available)
const (
	NavBack       = "Esc: 返回"
	NavExit       = "Ctrl+Q: 退出"
	NavUpDown     = "↑/↓: 滚动/选择"
	NavLeftRight  = "←/→: 移动"
	NavPgUpPgDown = "PgUp/PgDn: 翻页"
	NavHomeEnd    = "Home/End: 首/末"
	NavEnter      = "Enter: 继续"
	NavYn         = "Y/N: 接受/拒绝"
	NavCtrlC      = "Ctrl+C: 取消"
	NavCtrlS      = "Ctrl+S: 保存"
	NavCtrlR      = "Ctrl+R: 重置"
	NavCtrlH      = "Ctrl+H: 显示/隐藏"
	NavTab        = "Tab: 补全"
	NavSeparator  = " • "
)

// Welcome Screen constants
const (
	// Form interface implementation
	WelcomeFormTitle       = "欢迎使用 PentAGI"
	WelcomeFormDescription = "PentAGI 是一个自主渗透测试平台，利用 AI 技术执行全面的安全评估。"
	WelcomeFormName        = "欢迎"
	WelcomeFormOverview    = `系统检查项：
• 环境配置文件是否存在
• Docker API 可访问性与版本兼容性
• Worker 环境就绪状态
• 系统资源（CPU、内存、磁盘空间）
• 外部依赖所需的网络连通性

所有检查通过后，即可进入配置向导，设置 LLM 提供商、监控与安全工具。

安装程序会逐个组件引导你完成配置，并针对不同部署场景给出建议。`

	// Configuration status messages
	WelcomeConfigurationFailed = "⚠ 未通过的检查：%s"
	WelcomeConfigurationPassed = "✓ 所有系统检查均已通过"

	// Workflow steps
	WelcomeWorkflowTitle = "安装流程："
	WelcomeWorkflowStep1 = "1. 接受最终用户许可协议"
	WelcomeWorkflowStep2 = "2. 配置 LLM 提供商（OpenAI、Anthropic 等）"
	WelcomeWorkflowStep3 = "3. 设置集成组件（Langfuse、可观测性）"
	WelcomeWorkflowStep4 = "4. 配置安全选项"
	WelcomeWorkflowStep5 = "5. 部署并启动 PentAGI 服务"
	WelcomeSystemReady   = "✓ 系统就绪 —— 按 Enter 继续"
)

// Troubleshooting on welcome screen constants
const (
	TroubleshootTitle = "系统要求未满足"

	// Environment file issues
	TroubleshootEnvFileTitle = "缺少环境配置文件"
	TroubleshootEnvFileDesc  = "PentAGI 配置需要 .env 文件，但未找到该文件或文件不可读。"
	TroubleshootEnvFileFix   = `解决方法：
1. 在安装目录中将 .env.example 复制为 .env
2. 编辑 .env，至少配置一个 LLM 提供商的 API 密钥
3. 确保文件具有读取权限（chmod 644 .env）

快速修复：
cp .env.example .env && chmod 644 .env`

	// Write permissions
	TroubleshootWritePermTitle = "需要写入权限"
	TroubleshootWritePermDesc  = "安装程序需要对配置目录的写入权限，以保存设置并部署服务。"
	TroubleshootWritePermFix   = `解决方法：
1. 检查目录权限：ls -la
2. 授予写入权限：chmod 755 .
3. 或在具有写入权限的位置运行安装程序
4. 确保有足够的磁盘空间`

	// Docker not installed
	TroubleshootDockerNotInstalledTitle = "未安装 Docker"
	TroubleshootDockerNotInstalledDesc  = "本系统未安装 Docker。PentAGI 需要 Docker 来运行容器。"
	TroubleshootDockerNotInstalledFix   = `解决方法：
1. 安装 Docker Desktop：https://docs.docker.com/get-docker/
2. Linux 用户：按发行版对应的说明安装
3. 验证安装：docker --version
4. 确保 docker 命令在 PATH 中`

	// Docker not running
	TroubleshootDockerNotRunningTitle = "Docker 守护进程未运行"
	TroubleshootDockerNotRunningDesc  = "已安装 Docker，但守护进程未运行。Docker 服务必须处于活动状态。"
	TroubleshootDockerNotRunningFix   = `解决方法：
1. 启动 Docker Desktop（Windows/Mac）
2. Linux：sudo systemctl start docker
3. 检查状态：docker ps
4. 若使用 DOCKER_HOST，请确认远程守护进程可访问`

	// Docker permission issues
	TroubleshootDockerPermissionTitle = "Docker 权限被拒绝"
	TroubleshootDockerPermissionDesc  = "当前用户账号没有访问 Docker 的权限。这在 Linux 系统上很常见。"
	TroubleshootDockerPermissionFix   = `解决方法：
1. 将用户加入 docker 组：sudo usermod -aG docker $USER
2. 注销并重新登录使变更生效
3. 或使用 sudo 运行（不建议用于生产环境）
4. 验证：docker ps（应无需 sudo 即可执行）`

	// Generic Docker API issues
	TroubleshootDockerAPITitle = "Docker API 连接失败"
	TroubleshootDockerAPIDesc  = "无法建立与 Docker API 的连接。可能是配置或网络问题导致。"
	TroubleshootDockerAPIFix   = `解决方法：
1. 检查 DOCKER_HOST 环境变量
2. 确认 Docker 正在运行：docker version
3. 远程 Docker：确保网络连通
4. 若使用 TCP 连接，请检查防火墙设置
5. 可尝试：export DOCKER_HOST=unix:///var/run/docker.sock`

	// Docker version issues
	TroubleshootDockerVersionTitle = "Docker 版本过低"
	TroubleshootDockerVersionDesc  = "当前 Docker 版本不兼容。PentAGI 需要 Docker 20.0.0 或更高版本。"
	TroubleshootDockerVersionFix   = `解决方法：
1. 将 Docker 升级到 20.0.0 或更高版本
2. 访问 https://docs.docker.com/engine/install/

当前版本：%s
要求版本：20.0.0+`

	// Docker Compose issues
	TroubleshootComposeTitle = "未找到 Docker Compose"
	TroubleshootComposeDesc  = "需要 Docker Compose v2 插件，但 `docker compose` 不可用。"
	TroubleshootComposeFix   = `解决方法：
1. 安装或升级 Docker Desktop，或为 Docker Engine 安装 Docker Compose v2 插件
2. 确认插件可用：docker compose version
3. 若仅安装了旧版 docker-compose，请同时安装 Docker Compose v2 插件

PentAGI 执行的是 "docker compose"，因此仅有旧版 "docker-compose" 并不足够。
文档：https://docs.docker.com/compose/install/`

	// Docker Compose version issues
	TroubleshootComposeVersionTitle = "Docker Compose 版本过低"
	TroubleshootComposeVersionDesc  = "当前 `docker compose` 版本不兼容。PentAGI 需要 Docker Compose 1.25.0 或更高版本。"
	TroubleshootComposeVersionFix   = `当前版本：%s
要求版本：1.25.0+

解决方法：
1. 将 Docker Desktop 或 Docker Compose v2 插件升级到更高版本
2. 用以下命令验证结果：docker compose version

文档：https://docs.docker.com/compose/install/`

	// Worker environment issues
	TroubleshootWorkerTitle = "无法访问 Worker 的 Docker 环境"
	TroubleshootWorkerDesc  = "无法连接到用于 worker 容器的 Docker 环境。可能是远程或本地 Docker 配置问题。"
	TroubleshootWorkerFix   = `解决方法：
1. 使用远程 Docker 时，在运行安装程序前设置环境变量：
   export DOCKER_HOST=tcp://remote:2376
   export DOCKER_CERT_PATH=/path/to/certs
   export DOCKER_TLS_VERIFY=1
2. 验证连接：docker -H $DOCKER_HOST ps
3. 使用本地 Docker 时，请不要设置这些变量
4. 检查防火墙是否放通 Docker 端口（2375/2376）
5. 若使用 TLS，请确保证书有效`

	// CPU issues
	TroubleshootCPUTitle = "CPU 核心数不足"
	TroubleshootCPUDesc  = "PentAGI 至少需要 2 个 CPU 核心才能正常运行。"
	TroubleshootCPUFix   = `当前系统有 %d 个 CPU 核心，但至少需要 2 个。

虚拟机用户：
1. 在虚拟机设置中增加 CPU 分配
2. 确保宿主机有足够资源

Docker Desktop 用户：
设置 → 资源 → CPUs：设置为 2 或更多`

	// Memory issues
	TroubleshootMemoryTitle = "内存不足"
	TroubleshootMemoryDesc  = "可用内存不足以运行所选组件。"
	TroubleshootMemoryFix   = `内存需求：
• 基础系统：0.5 GB
• PentAGI 核心：+0.5 GB
• Langfuse（如启用）：+1.5 GB
• Observability（如启用）：+1.5 GB

共需：%.1f GB
可用：%.1f GB

解决方法：
1. 关闭不必要的应用程序
2. 提高 Docker 内存上限
3. 禁用可选组件（Langfuse/Observability）`

	// Disk space issues
	TroubleshootDiskTitle = "磁盘空间不足"
	TroubleshootDiskDesc  = "可用磁盘空间不足以完成安装和运行。"
	TroubleshootDiskFix   = `磁盘需求：
• 基础安装：至少 5 GB
• 启用组件后：10 GB + 每个组件 2 GB
• Worker 镜像：25 GB（含 6GB+ 的 Kali 镜像）

需要：%.1f GB
可用：%.1f GB

解决方法：
1. 释放磁盘空间
2. 为 Docker 使用外部存储
3. 清理未使用的 Docker 资源：
   docker system prune -a`

	// Network issues
	TroubleshootNetworkTitle = "网络连接失败"
	TroubleshootNetworkDesc  = "无法访问必需的外部服务，这会导致无法下载 Docker 镜像和更新。"
	TroubleshootNetworkFix   = `失败的检查项：
%s

解决方法：
1. 验证网络连接：ping docker.io
2. 检查 DNS 解析：nslookup docker.io
3. 如使用代理，请在运行安装程序前设置：
   export HTTP_PROXY=http://proxy:port
   export HTTPS_PROXY=http://proxy:port
4. 如需长期使用代理，请写入 .env：
   PROXY_URL=http://proxy:port
5. 检查防火墙是否放通出站 HTTPS（443 端口）
6. 若 DNS 解析失败，可尝试更换 DNS 服务器`

	// Generic hint at the bottom
	TroubleshootFixHint = "\n请解决上述问题后重新运行安装程序。"

	// Network failure messages (used in checker/helpers.go)
	NetworkFailureDNS        = "• docker.io 的 DNS 解析失败"
	NetworkFailureHTTPS      = "• 无法通过 HTTPS 访问外部服务"
	NetworkFailureDockerPull = "• 无法从镜像仓库拉取 Docker 镜像"
)

// System Checks constants
const (
	ChecksTitle               = "系统检查"
	ChecksWarningFailed       = "⚠ 部分检查未通过"
	CheckEnvironmentFile      = "环境配置文件"
	CheckWritePermissions     = "写入权限"
	CheckDockerAPI            = "Docker API"
	CheckDockerVersion        = "Docker 版本"
	CheckDockerCompose        = "Docker Compose"
	CheckDockerComposeVersion = "Docker Compose 版本"
	CheckWorkerEnvironment    = "工作容器环境"
	CheckSystemResources      = "系统资源"
	CheckNetworkConnectivity  = "网络连通性"
)

// EULA Screen constants
const (
	// Form interface implementation
	EULAFormDescription = "PentAGI 使用的法律条款与条件"
	EULAFormName        = "EULA"
	EULAFormOverview    = `请阅读并接受最终用户许可协议（EULA），以继续安装 PentAGI。

EULA 包含以下内容：
• 软件许可条款与使用权利
• 责任限制与担保声明
• 数据收集与隐私政策
• 合规要求与使用限制
• 支持与维护条款

您必须完整浏览整份文档并接受条款，才能继续安装流程。

可使用方向键、PgUp/PgDn 或 Home/End 键浏览文档。`

	// Error and status messages
	EULAErrorLoadingTitle     = "# 加载 EULA 出错\n\n加载 EULA 失败：%v"
	EULAContentFallback       = "# EULA 内容\n\n%s\n\n---\n\n*注意：Markdown 渲染失败：%v*"
	EULAConfigurationRead     = "✓ EULA 已阅读"
	EULAConfigurationAccepted = "✓ EULA 已接受"
	EULAConfigurationPending  = "⚠ EULA 尚未阅读"
	EULALoading               = "正在加载 EULA..."
	EULAProgress              = "进度：%d%%"
	EULAProgressComplete      = " • 已完成"
)

// Main Menu Screen constants
const (
	MainMenuTitle       = "PentAGI 配置"
	MainMenuDescription = "配置 PentAGI 的全部组件与设置"
	MainMenuName        = "主菜单"
	MainMenuOverview    = `欢迎使用 PentAGI 配置中心。

需要配置的核心组件：
• LLM Providers（模型提供商）—— 用于自主测试的 AI 语言模型
• Monitoring（监控）—— 可观测性与分析平台
• Tools（工具）—— 增强测试能力的附加功能
• System Settings（系统设置）—— 环境与部署选项

依次完成各个部分，即可完成 PentAGI 的安装配置。`

	MenuTitle        = "配置菜单"
	MenuSystemStatus = "系统状态"
)

// Main Menu Status Labels (not used)
const (
	MainMenuStatusPentagiRunning     = "PentAGI 已在运行"
	MainMenuStatusPentagiNotRunning  = "可以启动 PentAGI 服务"
	MainMenuStatusUpToDate           = "PentAGI 已是最新版本"
	MainMenuStatusUpdatesAvailable   = "有可用更新"
	MainMenuStatusReadyToStart       = "准备启动"
	MainMenuStatusAllServicesRunning = "所有服务均在运行"
	MainMenuStatusNoUpdatesAvailable = "暂无可用更新"
)

// LLM Providers Screen constants
const (
	LLMProvidersTitle       = "LLM 提供商配置"
	LLMProvidersDescription = "为 AI 智能体配置大语言模型提供商"
	LLMProvidersName        = "LLM 提供商"
	LLMProvidersOverview    = `PentAGI 使用多个专职 AI 智能体（研究员、开发者、执行器、渗透测试员），它们需要不同的 LLM 能力才能取得最佳渗透测试效果。

为什么要配置多个提供商：
• 智能体专职化：不同智能体分别受益于擅长推理、编码或分析的模型
• 成本优化：复杂任务使用高价推理模型（o3、grok-4、claude-sonnet-4、gemini-2.5-pro），简单操作使用低价模型
• 性能优化：各提供商各有所长 — OpenAI 适合中等任务，Anthropic 适合复杂任务，Gemini 更省成本

提供商选择建议：
• 云端生产环境：OpenAI + Anthropic + Gemini，性能与可靠性业界领先
• 企业/合规场景：AWS Bedrock，满足 SOC2、HIPAA 要求，并可访问多个模型家族
• 隐私/本地部署：Ollama 或 vLLM 搭配 Llama 3.1、Qwen3 等开源模型，数据完全自主可控

OpenRouter、DeepInfra、vLLM、Ollama 等提供商的开箱即用配置位于容器内的 /opt/pentagi/conf/ 目录`
)

// LLM Provider titles and descriptions
const (
	LLMProviderOpenAI        = "OpenAI"
	LLMProviderAnthropic     = "Anthropic"
	LLMProviderGemini        = "Google Gemini"
	LLMProviderBedrock       = "AWS Bedrock"
	LLMProviderOllama        = "Ollama"
	LLMProviderDeepSeek      = "DeepSeek"
	LLMProviderGLM           = "GLM Zhipu AI"
	LLMProviderKimi          = "Kimi Moonshot AI"
	LLMProviderQwen          = "Qwen Alibaba Cloud"
	LLMProviderCustom        = "Custom"
	LLMProviderOpenAIDesc    = "业界领先的 GPT 模型，综合表现出色"
	LLMProviderAnthropicDesc = "Claude 模型，推理能力与安全特性更强"
	LLMProviderGeminiDesc    = "Google 的先进多模态模型，知识覆盖面广"
	LLMProviderBedrockDesc   = "通过 AWS 企业级服务访问多家基础模型提供商"
	LLMProviderOllamaDesc    = "本地与云端开源模型，兼顾隐私与灵活性"
	LLMProviderDeepSeekDesc  = "国产先进模型，推理与多语言能力强"
	LLMProviderGLMDesc       = "智谱 AI 的 GLM 模型，擅长中英文任务"
	LLMProviderKimiDesc      = "月之暗面的长上下文模型，适合文档分析"
	LLMProviderQwenDesc      = "阿里云通义千问模型，适合多语言任务"
	LLMProviderCustomDesc    = "自定义 OpenAI 兼容端点，灵活性最高"
)

// Provider-specific help text
const (
	LLMFormOpenAIHelp = `OpenAI 提供业界领先的模型，其前沿推理能力非常适合复杂的渗透测试场景。

PentAGI 默认模型：
• o1、o4-mini：高级推理模型，用于复杂漏洞分析与策略规划
• GPT-4.1、GPT-4.1-mini：旗舰模型，针对漏洞利用开发与代码生成优化
• 根据智能体类型和任务复杂度自动选择模型

主要优势：
• 最先进的推理能力，支持逐步分析（o 系列模型）
• 出色的编码能力，适合自定义漏洞利用开发与载荷生成
• 性能可靠，服务稳定，API 文档完善
• 在安全研究与渗透测试场景中久经验证

适用场景：需要前沿 AI 能力的生产环境，以及性能优先于成本的团队
成本：定价偏高，但经过优化的配置能在成本与质量之间取得平衡

配置方式：在 https://platform.openai.com/api-keys 获取 API 密钥`

	LLMFormAnthropicHelp = `Anthropic Claude 模型在注重安全合规的渗透测试中表现优异，具备出色的推理与分析能力。

PentAGI 默认模型：
• Claude Sonnet-4：高端推理模型，用于复杂安全分析与策略性漏洞评估
• Claude 3.5 Haiku：高速模型，针对快速信息收集与简单解析任务优化
• 在各类安全测试场景中兼顾成本与性能

主要优势：
• 高度重视安全与伦理，在保持安全测试效果的同时减少有害输出
• 推理能力出色，适合系统化的漏洞分析与渗透测试方法
• 上下文窗口大，非常适合分析大型代码库与配置文件
• 善于理解复杂的安全场景与合规要求

适用场景：重视负责任测试实践的安全团队、注重合规的环境、需要深入分析的场景
成本：中等定价，在推理密集型安全工作流中性价比出色

配置方式：在 https://console.anthropic.com/ 获取 API 密钥`

	LLMFormGeminiHelp = `Google Gemini 兼具多模态能力与高级推理能力，非常适合全面的安全评估。

PentAGI 默认模型：
• Gemini 2.5 Pro：高级推理模型，用于深度漏洞分析和复杂漏洞利用开发
• Gemini 2.5 Flash：高性能模型，在速度与智能之间取得平衡，适用于大多数安全测试任务
• Gemini 2.0 Flash Lite：高性价比模型，用于快速扫描和信息收集
• 具备逐步分析的推理能力，可支撑深入的渗透测试

主要优势：
• 支持多模态，可分析截图、网络拓扑图和安全文档
• 价格有竞争力，开发和测试环境的速率限制较为宽松
• 上下文窗口大（最高 200 万 tokens），可分析超大代码库和系统配置
• 在多种编程语言的代码分析与漏洞识别方面表现出色

适用场景：预算敏感的团队、开发环境，以及需要分析图片/文档的场景
成本：主流云厂商中性价比最高，性能与价格比出色

配置方式：在 https://aistudio.google.com/app/apikey 获取 API 密钥`

	LLMFormBedrockHelp = `AWS Bedrock 提供企业级的 20 多种基础模型访问能力，支持多种认证方式并具备增强的安全特性。

PentAGI 默认模型：
• Claude Sonnet-4.5（通过 Bedrock）：高端推理模型，具备 AWS 企业级安全和扩展思考能力
• OpenAI GPT OSS 120B：推理能力强，适合科学分析和复杂安全任务
• Claude Haiku-4.5、DeepSeek V3.2、Qwen3-32B：高效模型，用于特定智能体角色和成本优化
• 通过统一接口访问 Amazon Nova（多模态）、Mistral、Moonshot 等更多模型

认证方式（按优先级）：
1. AWS 默认认证（BEDROCK_DEFAULT_AUTH=true）：使用 AWS SDK 凭证链，推荐用于 EC2/ECS/Lambda
2. Bearer 令牌（BEDROCK_BEARER_TOKEN）：基于令牌的认证，适用于自定义认证场景
3. 静态凭证（ACCESS_KEY + SECRET_KEY）：传统 IAM 凭证，适用于开发和测试

主要优势：
• 企业合规：具备 SOC2、HIPAA、FedRAMP 认证，支持数据驻留和治理控制
• 多提供商接入：来自 Anthropic、Amazon、OpenAI、Qwen、DeepSeek、Cohere、Mistral、Moonshot 的 20 多种模型
• 灵活认证：三种方式适配不同部署场景和安全要求
• 增强安全：VPC 集成、CloudTrail 日志、IAM 控制、私有端点实现完全隔离
• 区域化部署：可部署在指定 AWS 区域，优化延迟并满足数据主权要求

适用场景：企业环境、受监管行业，以及需要合规控制和灵活认证的团队
成本：价格有竞争力并提供预置吞吐量选项，但新账号的速率限制较严（每分钟 2-20 次请求）
重要提示：用于生产渗透测试流程时，请通过 AWS Service Quotas 控制台申请提升配额

配置方式：选择认证方式并配置凭证。速率限制详见 https://docs.aws.amazon.com/bedrock/`

	LLMFormOllamaHelp = `Ollama 支持两种部署方式，灵活性极高。

方式一：本地 Ollama 服务器（自托管）
• 在自有硬件上运行 Ollama（建议 8GB 以上内存，GPU 可选但有帮助）
• 数据完全私有 —— 所有处理都在本地完成
• 无持续费用 —— 仅需承担基础设施成本
• 无需 API 密钥 —— 通过网络访问控制来鉴权
• 配置方式：从 https://ollama.ai/ 安装，并设置 OLLAMA_SERVER_URL=http://ollama-server:11434

方式二：Ollama Cloud（托管服务）
• 云端托管模型，无需本地基础设施
• 无需硬件 —— 模型运行在 Ollama 的基础设施上
• 按量计费，并提供免费额度
• 需要 API 密钥 —— 在 https://ollama.com/settings/keys 生成
• 配置方式：在 https://ollama.com 注册，设置 OLLAMA_SERVER_URL=https://ollama.com 和 OLLAMA_SERVER_API_KEY=你的密钥

PentAGI 默认模型：
• Llama 3.1:8b、Qwen3:32b 等开源模型
• 可自由定制 —— 可在 100 多个可用模型间切换
• 支持模型自动下载与加载，使用便捷

主要优势：
• 双部署方式：可在隐私性（本地）与便捷性（云端）之间自由选择
• 成本灵活：本地部署无持续费用，云端按量计费
• 模型库丰富：可使用最新的开源模型（Llama、Qwen、Mistral、Gemma 等）
• 支持离网环境：本地部署可在隔离网络中运行

适用场景：注重隐私的团队（本地）、预算有限的部署（云端）、有数据主权要求的组织
配置选项：从 https://10.10.10.10:11434 本地安装，或在 https://ollama.com 注册云服务`

	LLMFormDeepSeekHelp = `DeepSeek 提供推理能力出色、多语言支持良好的先进 AI 模型。

PentAGI 默认模型：
• deepseek-v4-flash：性价比高的通用模型，适用于对话、代码生成和工具调用
• deepseek-v4-pro：更高阶的推理模型，适用于复杂逻辑、数学推理和安全分析
• 定价经济，性能可与主流模型竞争

核心优势：
• 编码与推理能力强，适合安全分析和漏洞利用开发
• 支持中英双语，适用于跨国渗透测试场景
• 定价有竞争力，性价比出色
• 兼容 OpenAI 接口，可无缝集成

LiteLLM 集成（多provider 统一网关）：
• 使用 LiteLLM 代理时，将 Provider Name 设为 'deepseek'
• 无需修改 config.yml 即可启用模型前缀（如 deepseek/deepseek-v4-flash）
• 直连 DeepSeek API 时此项为可选

适用场景：需要多语言支持的团队、注重成本的部署、中文安全测试
成本：定价极具竞争力，性能表现优秀

配置方式：在 https://platform.deepseek.com/ 获取 API 密钥`

	LLMFormGLMHelp = `GLM 由智谱 AI（源自清华大学）研发，具备出色的自然语言处理与推理能力。

PentAGI 默认模型：
• GLM-4-Air：高性能通用对话模型，针对常规任务和工具调用优化
• GLM-4-Plus：旗舰模型，推理与代码生成能力强
• GLM-Z1-Plus：高阶推理模型，具备面向安全研究的深度分析能力

核心优势：
• 中英文自然语言处理能力突出
• 在多语言安全测试与分析场景中表现优异
• GLM-4 与 GLM-Z1 系列在推理和编码方面均有增强
• 兼容 OpenAI 接口，易于集成

可选 API 端点：
• 国际站：https://api.z.ai/api/paas/v4（默认）
• 中国站：https://open.bigmodel.cn/api/paas/v4
• 编码专用：https://api.z.ai/api/coding/paas/v4

LiteLLM 集成（多 provider 统一网关）：
• 使用 LiteLLM 代理时，将 Provider Name 设为 'zai'
• 无需修改 config.yml 即可启用模型前缀（如 zai/glm-4）
• 直连 GLM API 时此项为可选

适用场景：中英双语渗透测试、在亚洲市场运营的团队
成本：定价有竞争力，多语言任务性能良好

配置方式：在 https://open.bigmodel.cn/ 获取 API 密钥`

	LLMFormKimiHelp = `Kimi（Moonshot AI 出品）提供超长上下文模型，非常适合分析大型代码库和文档。

PentAGI 默认模型：
• Moonshot-v1-8k：长上下文模型，支持最多 8K tokens，用于常规对话
• Kimi-k2.5：进阶模型，推理与文档理解能力强
• 针对处理大量文本和代码做过优化

主要优势：
• 超长上下文窗口（最多 1M tokens），可完整分析整个代码库
• 中英文支持出色，适合多语言渗透测试
• 文档密集型安全评估和威胁情报分析的性价比高
• 擅长理解复杂系统架构和长篇技术文档

可选 API 端点：
• 国际站：https://api.moonshot.ai/v1（默认）
• 中国站：https://api.moonshot.cn/v1

LiteLLM 集成：
• 使用 LiteLLM 代理时，将 Provider Name 设为 'moonshot'
• 无需修改 config.yml 即可启用模型前缀（如 moonshot/kimi-k2.5）
• 直接调用 Kimi API 时可不填

适用场景：大型代码库分析、文档密集型评估、安全研究中需要超长上下文的团队
成本：定价有竞争力，长上下文场景性价比突出

配置方式：在 https://platform.moonshot.ai/ 获取 API 密钥`

	LLMFormQwenHelp = `Qwen（阿里云百炼 / DashScope 出品）提供强大的多语言模型，并具备多模态能力。

PentAGI 默认模型：
• Qwen-Turbo：最快的轻量模型，适合高频任务和实时响应场景
• Qwen-Plus：性能均衡，适合常规对话、代码生成和工具调用
• Qwen-Max：旗舰推理模型，指令遵循能力强，可处理复杂任务
• QwQ-Plus：深度推理模型，支持长链思考，适合复杂逻辑分析

主要优势：
• 多语言支持出色（中文、英文及多种其他语言）
• 借助 Qwen-VL 具备多模态能力，可做视觉化安全分析
• 可与阿里云集成，便于企业级部署
• DashScope 生态提供额外的 AI 服务和工具
• Qwen2.5、Qwen3、QwQ 系列涵盖多种规模和专长

可选 API 端点：
• 美国：https://dashscope-us.aliyuncs.com/compatible-mode/v1（默认）
• 新加坡：https://dashscope-intl.aliyuncs.com/compatible-mode/v1
• 中国：https://dashscope.aliyuncs.com/compatible-mode/v1

LiteLLM 集成：
• 使用 LiteLLM 代理时，将 Provider Name 设为 'dashscope'
• 无需修改 config.yml 即可启用模型前缀（如 dashscope/qwen-plus）
• 直接调用 Qwen API 时可不填

适用场景：亚洲市场的团队、多语言安全测试、用 Qwen-VL 做视觉分析、已接入阿里云生态
成本：定价有竞争力，分层灵活，可匹配不同使用场景

配置方式：在 https://dashscope.console.aliyun.com/ 获取 API 密钥`

	LLMFormCustomHelp = `可配置任意兼容 OpenAI 协议的 API 端点，灵活性最高，便于对接既有基础设施。

开箱可用的配置：
• vLLM 部署：本地高吞吐推理，GPU 利用率最优
• OpenRouter：通过单一 API 访问 200 多个模型，价格有竞争力
• DeepInfra：面向主流开源模型的 Serverless 推理，按量计费
• Together AI、Groq、Fireworks：可选云服务商，各有专门的性能优化
• LiteLLM Proxy：通用网关，接入 100 多家服务商，支持负载均衡和统一接口（用 LLM_SERVER_PROVIDER 指定模型前缀）
• 部分推理模型和服务商在使用工具调用时需要保留推理内容（LLM_SERVER_PRESERVE_REASONING=true）

常见本地部署方案：
• vLLM：生产级服务，支持 Qwen、Llama、Mistral，具备批处理和 GPU 优化
• LocalAI：兼容 OpenAI 协议的封装层，支持多种本地模型和向量化服务
• Text Generation WebUI：社区热门界面，模型支持广泛，可做微调
• Hugging Face TGI：企业级文本生成推理，支持自动扩缩容和监控

主要优势：
• 灵活度无上限：可对接任意兼容 OpenAI 协议的端点或服务
• 成本优化：可选择价格更优的服务商，或在自有设施上部署模型
• 摆脱厂商锁定：可在服务商和模型之间无缝切换
• 自定义微调：可部署针对自身安全测试场景训练的专用模型

适用场景：对模型有特定要求、需要控制成本，或已有 LLM 基础设施的团队
LiteLLM 集成：将 LLM_SERVER_PROVIDER 设为对应的服务商名称（如 "openrouter"、"moonshot"），即可让同一套配置文件同时用于直连 API 和 LiteLLM 代理
配置示例：容器内 /opt/pentagi/conf/ 目录下提供了主流服务商的预置配置`
)

// LLM Provider Form field labels and descriptions
const (
	LLMFormFieldBaseURL           = "Base URL"
	LLMFormFieldAPIKey            = "API 密钥"
	LLMFormFieldDefaultAuth       = "使用 AWS 默认认证"
	LLMFormFieldBearerToken       = "Bearer Token"
	LLMFormFieldAccessKey         = "Access Key ID"
	LLMFormFieldSecretKey         = "Secret Access Key"
	LLMFormFieldSessionToken      = "Session Token"
	LLMFormFieldRegion            = "区域"
	LLMFormFieldModel             = "模型"
	LLMFormFieldConfigPath        = "配置文件路径"
	LLMFormFieldLegacyReasoning   = "旧版推理模式"
	LLMFormFieldPreserveReasoning = "保留推理内容"
	LLMFormFieldProviderName      = "提供商名称"
	LLMFormFieldPullTimeout       = "模型下载超时"
	LLMFormFieldPullEnabled       = "自动拉取模型"
	LLMFormFieldLoadModelsEnabled = "从服务端加载模型列表"
	LLMFormBaseURLDesc            = "该提供商的 API 端点地址"
	LLMFormAPIKeyDesc             = "用于身份验证的 API 密钥"
	LLMFormDefaultAuthDesc        = "使用 AWS SDK 默认凭证链（环境变量、EC2 角色、~/.aws/credentials），优先级最高"
	LLMFormBearerTokenDesc        = "用于身份验证的 Bearer Token，优先级高于静态凭证"
	LLMFormAccessKeyDesc          = "静态凭证认证所用的 AWS Access Key ID"
	LLMFormSecretKeyDesc          = "静态凭证认证所用的 AWS Secret Access Key"
	LLMFormSessionTokenDesc       = "临时凭证所用的 AWS Session Token（可选，与静态凭证配合使用）"
	LLMFormRegionDesc             = "Bedrock 服务所在的 AWS 区域"
	LLMFormModelDesc              = "该提供商默认使用的模型"
	LLMFormConfigPathDesc         = "配置文件路径（可选）"
	LLMFormLegacyReasoningDesc    = "启用旧版推理模式（true/false）"
	LLMFormPreserveReasoningDesc  = "在多轮对话中保留推理内容（部分提供商必需）"
	LLMFormProviderNameDesc       = "模型名称的提供商前缀（用于 LiteLLM 代理时有用）"
	LLMFormPullTimeoutDesc        = "下载模型的超时秒数（默认：600）"
	LLMFormPullEnabledDesc        = "启动时自动下载所需模型"
	LLMFormLoadModelsEnabledDesc  = "从 Ollama 服务端加载可用模型列表"
	LLMFormOllamaAPIKeyDesc       = "Ollama Cloud API 密钥（可选，使用本地 Ollama 服务端时留空）"
)

// LLM Provider Form status messages
const (
	LLMProviderFormTitle       = "LLM Provider %s Configuration"
	LLMProviderFormDescription = "Configure your Large Language Model provider settings"
	LLMProviderFormName        = "LLM Provider %s"
	LLMProviderFormOverview    = `Agent Role Assignment:
• Primary Agent & Pentester: Use reasoning models (o3, grok-4, claude-sonnet-4, gemini-2.5-pro) for complex vulnerability analysis
• Assistant & Adviser: Advanced models (o4-mini, claude-sonnet-4) for strategic planning and recommendations
• Coder & Installer: Precision models (gpt-4.1, claude-sonnet-4) for exploit development and system configuration
• Searcher & Enricher: Fast models (gpt-4.1-mini, claude-3.5-haiku, gemini-2.0-flash-lite) for information gathering
• Simple tasks: Lightweight models for JSON parsing and basic operations

Performance Considerations:
• Reasoning models provide step-by-step analysis but are slower and more expensive
• Standard models offer faster responses suitable for high-frequency agent interactions
• Each agent type uses provider-specific model configurations optimized for security testing workflows

Your configuration will determine which models each agent uses for different penetration testing scenarios.`
)

// Monitoring Screen
const (
	MonitoringTitle       = "Monitoring Configuration"
	MonitoringDescription = "Configure monitoring and observability platforms for comprehensive system insights"
	MonitoringName        = "Monitoring"
	MonitoringOverview    = `Comprehensive monitoring and observability for production-ready deployments.

Why monitoring matters:
• Track performance bottlenecks: Identify slow LLM calls, database queries, and system resources
• Debug issues faster: Detailed traces help diagnose problems across distributed components
• Optimize costs: Monitor token usage patterns and optimize expensive LLM interactions
• Production readiness: Essential for reliable operation in critical environments

Platform Options:
Langfuse: Specialized LLM observability with conversation tracking, prompt engineering insights, and cost analytics
Observability: Full-stack monitoring with metrics, traces, logs, and alerting for infrastructure and application health

Quick Setup:
• Development: Enable Langfuse for LLM insights only
• Production: Enable both platforms for comprehensive monitoring
• Cost-conscious: Use embedded modes to avoid external service fees`
)

// Langfuse Integration constants
const (
	MonitoringLangfuseFormTitle       = "Langfuse Configuration"
	MonitoringLangfuseFormDescription = "Configuration of Langfuse integration for LLM monitoring"
	MonitoringLangfuseFormName        = "Langfuse"
	MonitoringLangfuseFormOverview    = `Langfuse provides:
• Complete conversation tracking
• Model performance metrics
• Cost monitoring and optimization
• User behavior analytics
• Debug traces for AI interactions

Choose between embedded instance or external connection.`

	// Deployment types
	MonitoringLangfuseEmbedded = "内置服务"
	MonitoringLangfuseExternal = "外部服务"
	MonitoringLangfuseDisabled = "已禁用"

	// Form fields
	MonitoringLangfuseDeploymentType     = "部署方式"
	MonitoringLangfuseDeploymentTypeDesc = "选择 Langfuse 的部署方式"
	MonitoringLangfuseBaseURL            = "服务地址"
	MonitoringLangfuseBaseURLDesc        = "Langfuse 服务器地址（例如 https://cloud.langfuse.com）"
	MonitoringLangfuseProjectID          = "项目 ID"
	MonitoringLangfuseProjectIDDesc      = "Langfuse 中的项目标识符"
	MonitoringLangfusePublicKey          = "公钥"
	MonitoringLangfusePublicKeyDesc      = "用于访问项目的公开 API 密钥"
	MonitoringLangfuseSecretKey          = "私钥"
	MonitoringLangfuseSecretKeyDesc      = "用于访问项目的私密 API 密钥"
	MonitoringLangfuseListenIP           = "监听 IP"
	MonitoringLangfuseListenIPDesc       = "Docker 端口映射使用的绑定地址（例如 0.0.0.0 表示监听所有网卡）"
	MonitoringLangfuseListenPort         = "监听端口"
	MonitoringLangfuseListenPortDesc     = "Docker 为 Langfuse Web 界面暴露的外部 TCP 端口"

	// Admin settings for embedded
	MonitoringLangfuseAdminEmail        = "管理员邮箱"
	MonitoringLangfuseAdminEmailDesc    = "用于登录 Langfuse 管理后台的邮箱"
	MonitoringLangfuseAdminPassword     = "管理员密码"
	MonitoringLangfuseAdminPasswordDesc = "用于登录 Langfuse 管理后台的密码"
	MonitoringLangfuseAdminName         = "管理员用户名"
	MonitoringLangfuseAdminNameDesc     = "Langfuse 中的管理员用户名"
	MonitoringLangfuseLicenseKey        = "企业版许可证密钥"
	MonitoringLangfuseLicenseKeyDesc    = "Langfuse 企业版许可证密钥（可选）"

	// Help text
	MonitoringLangfuseModeGuide    = "选择部署方式：内置（本地自主管理）、外部（云端/已有实例）、禁用（不启用分析）"
	MonitoringLangfuseEmbeddedHelp = `内置模式会部署完整的 Langfuse 组件栈：
• PostgreSQL + ClickHouse 数据库
• MinIO S3 存储 + Redis 缓存
• 完整的 LLM 会话追踪
• 成本分析与性能指标
• 数据全部保留在你自己的服务器上

资源需求：
• 至少约 2GB 内存、5GB 磁盘空间
• 会话日志需要额外存储空间
• 自动完成部署与维护

适用场景：注重数据隐私、需要自定义配置或不希望依赖外部服务的团队。所有分析数据存储在本地，并拥有完整的管理权限。

默认管理员访问方式：
• Web 界面：http://localhost:4000
• 登录账号：admin@pentagi.com
• 密码：password（首次登录需修改）`
	MonitoringLangfuseExternalHelp = `外部模式连接 cloud.langfuse.com 或你已有的 Langfuse 服务器：

• 无需本地基础设施
• 由服务方负责升级与维护
• 可在团队间共享分析数据
• 可使用企业版功能
• 数据存储在外部服务方

配置前提：
• 拥有 Langfuse 账号及 API 密钥
• 需要可用的互联网连接
• 需要项目 ID 与认证密钥

适用场景：使用云服务、希望免运维，或需要与组织内已有 Langfuse 部署集成的团队。`
	MonitoringLangfuseDisabledHelp = `Langfuse 已禁用。缺少 LLM 可观测性后，你将无法获得：

• 会话历史追踪
• Token 用量与成本分析
• 模型性能指标
• AI 交互的调试链路
• 用户行为分析
• 提示词优化洞察

建议在生产环境中启用，
以便监控 AI 智能体表现
并有效优化成本。`
)

// Graphiti Integration constants
const (
	MonitoringGraphitiFormTitle       = "Graphiti 配置（测试版）"
	MonitoringGraphitiFormDescription = "配置 Graphiti 知识图谱集成"
	MonitoringGraphitiFormName        = "Graphiti（测试版）"
	MonitoringGraphitiFormOverview    = `⚠️  测试功能：该功能仍在积极开发中，请关注后续更新以获取功能改进与稳定性修复。

Graphiti 提供时序知识图谱能力：
• 实体与关系抽取
• 面向 AI 智能体的语义记忆
• 时序上下文追踪
• 跨任务流复用已有知识

⚠️  前置要求：Graphiti 需要已配置的 OpenAI 提供商（LLM 提供商 → OpenAI）来完成实体抽取。

请选择使用内置实例还是连接外部服务。`

	// Deployment types
	MonitoringGraphitiEmbedded = "内置组件栈"
	MonitoringGraphitiExternal = "外部服务"
	MonitoringGraphitiDisabled = "已禁用"

	// Form fields
	MonitoringGraphitiDeploymentType     = "部署方式"
	MonitoringGraphitiDeploymentTypeDesc = "选择 Graphiti 的部署方式"
	MonitoringGraphitiURL                = "Graphiti 服务地址"
	MonitoringGraphitiURLDesc            = "Graphiti API 服务器地址"
	MonitoringGraphitiTimeout            = "请求超时"
	MonitoringGraphitiTimeoutDesc        = "Graphiti 操作的超时时间（秒）"
	MonitoringGraphitiModelName          = "抽取模型"
	MonitoringGraphitiModelNameDesc      = "用于实体抽取的 LLM 模型（使用 LLM 提供商配置中的 OpenAI 提供商）"
	MonitoringGraphitiNeo4jUser          = "Neo4j 用户名"
	MonitoringGraphitiNeo4jUserDesc      = "访问 Neo4j 数据库的用户名"
	MonitoringGraphitiNeo4jPassword      = "Neo4j 密码"
	MonitoringGraphitiNeo4jPasswordDesc  = "访问 Neo4j 数据库的密码"
	MonitoringGraphitiNeo4jDatabase      = "Neo4j 数据库"
	MonitoringGraphitiNeo4jDatabaseDesc  = "Neo4j 数据库名称"

	// Help text
	MonitoringGraphitiModeGuide    = "选择部署方式：内置（本地 Neo4j）、外部（已有 Graphiti）、禁用（不启用知识图谱）"
	MonitoringGraphitiEmbeddedHelp = `⚠️  BETA: This feature is under active development. Monitor updates for improvements.

Embedded deploys complete Graphiti stack:
• Neo4j graph database
• Graphiti API service
• Automatic entity extraction from agent interactions
• Temporal relationship tracking
• Private knowledge graph on your server

Prerequisites:
• OpenAI provider must be configured (LLM Providers → OpenAI)
• OpenAI API key is used for entity extraction
• Configured model will be used for knowledge graph operations

Resource requirements:
• ~1.5GB RAM, 3GB disk space minimum
• Neo4j UI: http://localhost:7474
• Graphiti API: http://localhost:8000
• Automatic setup and maintenance

Best for: Teams wanting knowledge graph capabilities with full data control and privacy.`
	MonitoringGraphitiExternalHelp = `⚠️  BETA: This feature is under active development. Monitor updates for improvements.

External connects to your existing Graphiti server:

• No local infrastructure needed
• Managed updates and maintenance
• Shared knowledge graph across teams
• Data stored on external provider

Setup requirements:
• Graphiti server URL and access
• Network connectivity required
• External server must be configured with OpenAI API key
• Model and extraction settings configured on external server

Best for: Teams using existing Graphiti deployments or cloud services.`
	MonitoringGraphitiDisabledHelp = `Graphiti is disabled. You will not have:

• Temporal knowledge graph
• Entity and relationship extraction
• Semantic memory for AI agents
• Knowledge reuse across flows
• Advanced contextual search

Note: Graphiti is currently in beta.
Consider enabling for production use
to build a knowledge base from
penetration testing results.`
)

// Observability Integration constants
const (
	MonitoringObservabilityFormTitle       = "Observability Configuration"
	MonitoringObservabilityFormDescription = "Configuration of monitoring and observability stack"
	MonitoringObservabilityFormName        = "Observability"
	MonitoringObservabilityFormOverview    = `Observability stack includes:
• Grafana dashboards for visualization
• VictoriaMetrics for time-series data
• Jaeger for distributed tracing
• Loki for log aggregation
• OpenTelemetry for data collection

Monitor PentAGI performance and system health.`

	// Deployment types
	MonitoringObservabilityEmbedded = "Embedded Stack"
	MonitoringObservabilityExternal = "External Collector"
	MonitoringObservabilityDisabled = "Disabled"

	// Form fields
	MonitoringObservabilityDeploymentType     = "Deployment Type"
	MonitoringObservabilityDeploymentTypeDesc = "Select the deployment type for monitoring"
	MonitoringObservabilityOTelHost           = "OpenTelemetry Host"
	MonitoringObservabilityOTelHostDesc       = "Address of the external OpenTelemetry collector"

	// embedded listen fields
	MonitoringObservabilityGrafanaListenIP        = "Grafana Listen IP"
	MonitoringObservabilityGrafanaListenIPDesc    = "Bind address used by Docker port mapping (e.g., 0.0.0.0 to expose on all interfaces)"
	MonitoringObservabilityGrafanaListenPort      = "Grafana Listen Port"
	MonitoringObservabilityGrafanaListenPortDesc  = "External TCP port exposed by Docker for Grafana web UI"
	MonitoringObservabilityOTelGrpcListenIP       = "OTel gRPC Listen IP"
	MonitoringObservabilityOTelGrpcListenIPDesc   = "Bind address used by Docker port mapping (e.g., 0.0.0.0 to expose on all interfaces)"
	MonitoringObservabilityOTelGrpcListenPort     = "OTel gRPC Listen Port"
	MonitoringObservabilityOTelGrpcListenPortDesc = "External TCP port exposed by Docker for OTel gRPC receiver"
	MonitoringObservabilityOTelHttpListenIP       = "OTel HTTP Listen IP"
	MonitoringObservabilityOTelHttpListenIPDesc   = "Bind address used by Docker port mapping (e.g., 0.0.0.0 to expose on all interfaces)"
	MonitoringObservabilityOTelHttpListenPort     = "OTel HTTP Listen Port"
	MonitoringObservabilityOTelHttpListenPortDesc = "External TCP port exposed by Docker for OTel HTTP receiver"

	// Help text
	MonitoringObservabilityModeGuide    = "Choose monitoring: Embedded (full stack), External (existing infra), Disabled (no monitoring)"
	MonitoringObservabilityEmbeddedHelp = `Embedded deploys complete monitoring:
• Grafana dashboards and alerting
• VictoriaMetrics time-series database
• Jaeger distributed tracing UI
• Loki log aggregation system
• ClickHouse analytical database
• Node Exporter + cAdvisor metrics
• OpenTelemetry data collection

Auto-instrumented components with
pre-built dashboards for system health,
performance analysis, and debugging.

Resource requirements:
• ~1.5GB RAM, 3GB disk space minimum
• Grafana UI: http://localhost:3000
• Profiling: http://localhost:7777

Best for: Complete system visibility,
troubleshooting, and performance tuning.`
	MonitoringObservabilityExternalHelp = `External sends telemetry to your existing monitoring infrastructure:

• OTLP protocol over HTTP/2 (no TLS)
• Your collector must support:
  - OTLP HTTP receiver (port 4318)
  - OTLP gRPC receiver (port 8148)
  - tls: insecure: true setting
• Sends metrics, traces, and logs
• Compatible with enterprise platforms:
  Datadog, New Relic, Splunk, etc.

OTEL_HOST example:
your-collector:4318

Collector config requirement:
tls: insecure: true

Best for: Organizations with existing
monitoring infrastructure or centralized
observability platforms.`
	MonitoringObservabilityDisabledHelp = `Observability is disabled. You will not have:

• System performance monitoring
• Distributed request tracing
• Structured log aggregation
• Resource usage analytics
• Error tracking and alerting
• Performance bottleneck analysis

Consider enabling for production use
to monitor system health, debug issues,
and optimize performance effectively.`
)

// Summarizer Screen
const (
	SummarizerTitle       = "摘要器配置"
	SummarizerDescription = "启用对话摘要，以降低 LLM 成本并改善上下文管理"
	SummarizerName        = "摘要器"
	SummarizerOverview    = `优化上下文占用、降低 LLM 成本，并与模型能力相匹配。

何时需要调整摘要设置：
• Token 成本偏高：缩减上下文规模（4K-12K，而非 22K+ tokens）
• 出现"上下文过长"错误：按模型的上限进行配置
• 对话连贯性差：提升上下文保留量以改善质量
• 模型类型不同：针对短上下文与长上下文模型分别调优

General（通用摘要）：成本控制力最强、参数可精细调节，适合研究与分析类任务
Assistant（助手摘要）：对话质量最佳，具备智能上下文管理，适合交互式会话

快速见效的做法：
• 降低成本：使用 General，将"保留最近分段数"减至 1-2
• 上下文报错：将各项上限与你的模型对齐（8K/32K/128K）
• 优先保质量：使用 Assistant 并适当提高上限`

	SummarizerTypeGeneralName = "通用摘要（General）"
	SummarizerTypeGeneralDesc = "用于对话上下文管理的全局摘要设置"

	SummarizerTypeGeneralInfo = `适用于追求最强成本控制、并需兼容短上下文模型的场景。

在以下情况下最为合适：
• 需要大幅降低成本：可精调每个参数，将 token 用量压到最低
• 短上下文模型（8K-32K）：通过精确上限避免溢出错误
• 研究/分析类任务：在不丢失关键数据的前提下受控压缩
• 自定义问答处理：完全掌控问答对的处理方式

典型效果：
• 相比默认设置可降低 40-70% 成本
• 上下文规模 4K-12K tokens（Assistant 模式为 22K+）
• 在 GPT-3.5、Claude Instant 等较小模型上表现更好
• 可精确权衡"保留对话记忆"与"重置上下文"

最佳实践：
• 从 1-2 个"保留最近分段数"开始，以获得最大成本节省
• 启用"尺寸管理"以自动防止溢出
• 仅在关键推理任务中才关闭问答压缩`

	SummarizerTypeAssistantName = "助手摘要（Assistant）"
	SummarizerTypeAssistantDesc = "面向 AI 助手场景的专用摘要设置"

	SummarizerTypeAssistantInfo = `适用于追求最佳对话质量与连贯性的场景。

在以下情况下最为合适：
• 长推理链路：为复杂的多步思考维持上下文
• 高质量对话：保留对话流程与助手的表达风格
• 长上下文模型（64K+）：高效发挥模型的完整能力
• 交互式会话：更好地记住用户偏好与历史对话

典型效果：
• 上下文规模 8K-40K tokens，并具备智能压缩
• 对话连贯性显著优于手工设置
• 针对推理任务自动优化上下文
• 成本与质量兼顾（上下文约为 General 模式的 3 倍）

最佳实践：
• 多数场景直接使用默认值即可，它们已经过预优化
• 仅在任务非常复杂时才增加"保留最近分段数"
• 关注上下文用量，成本随 token 数线性增长
• 非常适合 GPT-4、Claude 等大上下文模型`
)

// Summarizer Form Screen
const (
	SummarizerFormGeneralTitle   = "通用摘要器配置"
	SummarizerFormAssistantTitle = "助手摘要器配置"
	SummarizerFormDescription    = "配置 %s 设置"

	// Field Labels and Descriptions
	SummarizerFormPreserveLast     = "尺寸管理"
	SummarizerFormPreserveLastDesc = "控制最后一个分段的压缩行为。启用：分段大小受 LastSecBytes 限制（上下文更小）。禁用：分段自由增长（上下文更大）"

	SummarizerFormUseQA     = "问答摘要"
	SummarizerFormUseQADesc = "当问答内容总量超过 MaxQABytes 或 MaxQASections 上限时，启用问答对压缩"

	SummarizerFormSumHumanInQA     = "压缩用户消息"
	SummarizerFormSumHumanInQADesc = "将用户消息一并纳入问答压缩。禁用：保留用户原文（多数场景推荐）"

	SummarizerFormLastSecBytes     = "分段大小上限"
	SummarizerFormLastSecBytesDesc = "启用尺寸管理时，每个最近分段的最大字节数。值越大：单个分段细节越多，token 用量越高"

	SummarizerFormMaxBPBytes     = "回复大小上限"
	SummarizerFormMaxBPBytesDesc = "单条 AI 回复在被压缩前的最大字节数。避免单条超长回复占据整个上下文"

	SummarizerFormMaxQASections     = "问答分段数上限"
	SummarizerFormMaxQASectionsDesc = "触发问答压缩前允许的最大问答分段数。与 MaxQABytes 配合，共同控制问答记忆总量"

	SummarizerFormMaxQABytes     = "问答记忆总量"
	SummarizerFormMaxQABytesDesc = "所有问答分段合计的最大字节数。超出后（与 MaxQASections 共同判定）触发问答压缩以压回上限内"

	SummarizerFormKeepQASections     = "保留最近分段数"
	SummarizerFormKeepQASectionsDesc = "不做压缩、原样保留的最近对话分段数量。影响上下文规模的首要参数"

	// Enhanced Help Text - General (common principles)
	SummarizerFormGeneralHelp = `上下文规模估算：常见 4K-22K tokens，参数拉满时可达 94K。

关键影响因素：
• 保留最近分段数：影响最大，每增加 1 个约增加 1.5-9K tokens
• 关闭尺寸管理：上下文增大 2-3 倍（压缩更少）
• 分段/回复上限：控制单个组成部分的大小
• 问答记忆：在超出上限时管理整体对话历史

参数之间的相互作用：
• 只有 MaxQABytes 与 MaxQASections 同时超限，才会触发问答压缩
• 禁用尺寸管理后，分段可增长到上限的 2 倍
• 回复上限可避免单条超长输出占满上下文
• 压缩用户消息（SummHumanInQA）约节省 5%，但会丢失原始措辞

小上下文模型建议下调：
• 保留最近分段数：1-2（默认为 3 以上）
• 分段上限：25-35KB（默认 50KB 以上）
• 简单对话可禁用尺寸管理

常见误区：
• 保留最近分段数设置过高（上下文溢出的主要原因）
• 启用尺寸管理的同时把分段上限设得过低（过度压缩）
• 问答上限不匹配（字节数高但分段数低，等于无效）

当前算法会压缩较早的内容，同时保持最近上下文的质量。`

	// Enhanced Help Text - Assistant specific (interactive conversations)
	SummarizerFormAssistantHelp = `Optimized for interactive conversations requiring context continuity.

Default tuning (3 Recent Sections, 75KB limits):
• Typical range: 8K-40K tokens
• Good for: Extended dialogues, reasoning chains, context-dependent tasks
• Models: Works well with 32K+ context models

Adjustments by model type:
• Short context (≤16K): Recent Sections=1-2, Section Limit=45KB
• Long context (128K+): Can increase Recent Sections=5-7
• High-frequency chat: Reduce Recent Sections=2 for faster responses

Advanced tuning:
• QA Memory 200KB+ for document analysis conversations
• Response Limit 24-32KB for detailed technical responses
• Keep User Messages uncompressed (SummHumanInQA=false) for better context

Performance optimization:
• Each Recent Section ≈ 9-18KB in assistant mode
• Size Management reduces growth by ~20% but may lose detail
• QA compression triggers less often due to larger default limits

Size Management enabled by default - maintains conversation flow while preventing context overflow.
Monitor actual token usage and adjust Recent Sections first, then limits.`

	// Context size estimation
	SummarizerContextEstimatedSize    = "Estimated context size: %s\n%s"
	SummarizerContextTokenRange       = "~%s tokens"
	SummarizerContextTokenRangeMinMax = "~%s-%s tokens"
	SummarizerContextRequires256K     = "Requires 256K+ context model"
	SummarizerContextRequires128K     = "Requires 128K+ context model"
	SummarizerContextRequires64K      = "Requires 64K+ context model"
	SummarizerContextRequires32K      = "Requires 32K+ context model"
	SummarizerContextRequires16K      = "Requires 16K+ context model"
	SummarizerContextFitsIn8K         = "Fits in 8K+ context model"
)

// Tools screen strings
const (
	ToolsTitle       = "工具配置"
	ToolsDescription = "通过附加工具和选项增强智能体能力"
	ToolsName        = "工具"
	ToolsOverview    = `为 AI 智能体配置附加工具与能力。
每个工具都可按需启用并单独配置。

可用设置：
• Human-in-the-loop（人工介入）- 允许测试过程中由用户参与决策
• AI Agents Settings（智能体设置）- 配置 AI 智能体的全局行为
• Search Engines（搜索引擎）- 配置外部搜索服务商
• Scraper（网页抓取）- 网页内容提取与分析
• Graphiti (beta)（知识图谱）- 用于语义记忆的时序知识图谱
• Docker（容器环境）- 容器运行环境配置`
)

// Server Settings screen strings
const (
	ServerSettingsFormTitle       = "服务器设置"
	ServerSettingsFormDescription = "配置 PentAGI 服务器的网络访问与对外路由"
	ServerSettingsFormName        = "服务器设置"
	ServerSettingsFormOverview    = `• 网络绑定 - 控制 PentAGI 监听的网卡与端口
• 公开 URL - 用于重定向的外部地址及可选的基础路径
• CORS - 允许通过浏览器访问的来源
• 代理 - 访问 LLM/搜索服务商时使用的 HTTP/HTTPS 代理
• SSL 目录 - 存放 server.crt 与 server.key（PEM 格式）的自定义证书目录
• 数据目录 - 智能体产物与任务流工作区的持久化存储`

	// Field labels and descriptions
	ServerSettingsLicenseKey     = "许可证密钥"
	ServerSettingsLicenseKeyDesc = "PentAGI 许可证密钥，格式为 XXXX-XXXX-XXXX-XXXX"

	ServerSettingsHost     = "服务器主机（监听 IP）"
	ServerSettingsHostDesc = "Docker 端口映射使用的绑定地址（例如 0.0.0.0 表示在所有网卡上暴露）"

	ServerSettingsPort     = "服务器端口（监听端口）"
	ServerSettingsPortDesc = "Docker 为 PentAGI Web 界面对外暴露的 TCP 端口"

	ServerSettingsPublicURL     = "公开 URL"
	ServerSettingsPublicURLDesc = "用于重定向和链接的公开基础 URL（支持基础路径，例如 https://example.com/pentagi/）"

	ServerSettingsCORSOrigins     = "CORS 允许来源"
	ServerSettingsCORSOriginsDesc = "以逗号分隔的允许来源列表（例如 https://localhost:8443,https://localhost）"

	ServerSettingsProxyURL     = "HTTP/HTTPS 代理"
	ServerSettingsProxyURLDesc = "访问 LLM 和外部工具时使用的代理（不用于 Docker API 访问）"

	ServerSettingsProxyUsername     = "代理用户名"
	ServerSettingsProxyUsernameDesc = "代理认证用户名（可选）"
	ServerSettingsProxyPassword     = "代理密码"
	ServerSettingsProxyPasswordDesc = "代理认证密码（可选）"

	ServerSettingsHTTPClientTimeout       = "HTTP 客户端超时"
	ServerSettingsHTTPClientTimeoutDesc   = "调用外部 API（LLM 服务商、搜索引擎等）的超时秒数"
	ServerSettingsTerminalToolTimeout     = "终端工具超时"
	ServerSettingsTerminalToolTimeoutDesc = "终端命令的默认超时秒数（填 0 或负数表示使用 3 小时上限）"

	ServerSettingsExternalSSLCAPath     = "自定义 CA 证书路径"
	ServerSettingsExternalSSLCAPathDesc = "容器内自定义根 CA 证书的路径（例如 /opt/pentagi/ssl/ca-bundle.pem）"

	ServerSettingsExternalSSLInsecure     = "跳过 SSL 校验"
	ServerSettingsExternalSSLInsecureDesc = "禁用 SSL/TLS 证书校验（仅用于自签名证书的测试场景）"

	ServerSettingsSSLDir     = "SSL 目录"
	ServerSettingsSSLDirDesc = "存放 PEM 格式 server.crt 与 server.key 的目录（server.crt 可包含完整证书链）"

	ServerSettingsDataDir     = "数据目录"
	ServerSettingsDataDirDesc = "存放所有智能体生成文件的目录；其中的 flow-N 子目录会挂载为工作容器内的 /work"

	ServerSettingsCookieSigningSalt     = "Cookie 签名盐值"
	ServerSettingsCookieSigningSaltDesc = "用于签名 Cookie 的密钥（请妥善保密）"

	// Hints for fields overview
	ServerSettingsLicenseKeyHint          = "许可证密钥"
	ServerSettingsHostHint                = "监听 IP"
	ServerSettingsPortHint                = "监听端口"
	ServerSettingsPublicURLHint           = "公开 URL"
	ServerSettingsCORSOriginsHint         = "CORS 来源"
	ServerSettingsProxyURLHint            = "代理地址"
	ServerSettingsProxyUsernameHint       = "代理用户名"
	ServerSettingsProxyPasswordHint       = "代理密码"
	ServerSettingsHTTPClientTimeoutHint   = "HTTP 超时"
	ServerSettingsTerminalToolTimeoutHint = "终端超时"
	ServerSettingsExternalSSLCAPathHint   = "自定义 CA 路径"
	ServerSettingsExternalSSLInsecureHint = "跳过 SSL 校验"
	ServerSettingsSSLDirHint              = "SSL 目录"
	ServerSettingsDataDirHint             = "数据目录"

	// Help texts per-field
	ServerSettingsGeneralHelp = `PentAGI exposes its web UI via Docker with configurable host and port.

Public URL must reflect how users reach the server. If using a subpath (e.g., /pentagi/), include it here. CORS controls browser access from specified origins. Proxy affects outbound traffic to LLM/search providers and other external services used by Tools.

SSL directory allows providing custom certificates. When set, server will use server.crt and server.key from that directory. Data directory stores artifacts and working files for flows.`

	ServerSettingsLicenseKeyHelp = `PentAGI License Key in format of XXXX-XXXX-XXXX-XXXX. It's used to communicate with PentAGI Cloud API.`

	ServerSettingsHostHelp = `Bind address for published port in docker-compose mapping.

Examples:
• 127.0.0.1 — local-only access
• 0.0.0.0 — expose on all interfaces`

	ServerSettingsPortHelp = `External port for PentAGI UI. Must be available on the host. Example: 8443.`

	ServerSettingsPublicURLHelp = `Set the public base URL used in redirects and links.

Examples:
• http://localhost:8443
• https://example.com/
• https://example.com/pentagi/ (with base path)`

	ServerSettingsCORSOriginsHelp = `Comma-separated allowed origins for browser access.`

	ServerSettingsProxyURLHelp = `HTTP or HTTPS proxy for outbound requests to LLM providers and external tools. Not used for Docker API communication.`

	ServerSettingsHTTPClientTimeoutHelp = `Timeout in seconds for all external HTTP/HTTPS API calls including:
• LLM provider requests (OpenAI, Anthropic, Bedrock, etc.)
• Search engine queries (Google, Tavily, Perplexity, etc.)
• External tool integrations
• Embedding generation requests

Default: 600 seconds (10 minutes)
Setting to 0 disables timeout (not recommended in production)
Too low values may cause legitimate long-running requests to fail.`

	ServerSettingsTerminalToolTimeoutHelp = `Default timeout in seconds applied when an agent requests timeout=0 or a negative timeout value.

This affects commands executed through the isolated terminal container, including scanners and CLI-based utilities.

Default: 1200 seconds (20 minutes)
Allowed range: 1–10800 seconds (up to 3 hours)
Values <= 0 or above 10800 are clamped to the maximum (10800 s = 3 hours); agents are never allowed to run indefinitely.
Explicit timeout values provided by the tool call override this default when they are within the 1–10800 s range.`

	ServerSettingsExternalSSLCAPathHelp = `Path to custom CA certificate file (PEM format) inside the container.

Must point to /opt/pentagi/ssl/ directory, which is mounted from pentagi-ssl volume on the host.

Examples:
• /opt/pentagi/ssl/ca-bundle.pem
• /opt/pentagi/ssl/corporate-ca.pem

File can contain multiple root and intermediate certificates.`

	ServerSettingsExternalSSLInsecureHelp = `Disable SSL/TLS certificate validation for connections to LLM providers and external services.

⚠ WARNING: Use only for testing with self-signed certificates. Never enable in production.

When enabled, all certificate validation is bypassed, making connections vulnerable to man-in-the-middle attacks.`

	ServerSettingsSSLDirHelp = `Path to directory with server.crt and server.key in PEM format. server.crt may include fullchain. Overrides default generated certificate behavior.`

	ServerSettingsDataDirHelp = `Host directory for persistent data. PentAGI stores agent artifacts under flow-N subdirectories, which map to /work inside worker containers.`

	ServerSettingsCookieSigningSaltHelp = `Secret salt used to sign cookies. Keep it private.`
)

// Human-in-the-loop screen strings
const (
	// AI Agents Settings screen strings
	ToolsAIAgentsSettingsFormTitle       = "AI 智能体设置"
	ToolsAIAgentsSettingsFormDescription = "配置 AI 智能体的全局行为"
	ToolsAIAgentsSettingsFormName        = "AI 智能体设置"
	ToolsAIAgentsSettingsFormOverview    = `本节配置 PentAGI 中 AI 智能体的全局行为。

基础设置：
• 启用用户交互：允许智能体在需要时请求用户输入
• 启用多智能体模式：允许助手调度多个专职智能体

执行监控（⚠️  测试版）：
• 启用执行监控：由导师智能体自动监督并分析执行模式
• 相同工具调用阈值：触发导师复查前允许的连续相同工具调用次数
• 累计工具调用阈值：触发导师复查前允许的累计工具调用次数

工具调用上限：
• 最大工具调用数（通用智能体）：防止 Assistant、Primary Agent、Pentester、Coder、Installer 失控执行
• 最大工具调用数（受限智能体）：防止 Searcher、Enricher、Memorist 等失控执行

任务规划（⚠️  测试版）：
• 启用任务规划：为专职智能体生成结构化执行计划

⚠️  测试版功能仍在积极开发中，建议仅用于测试。`

	// field labels and descriptions
	ToolsAIAgentsSettingHumanInTheLoop          = "启用用户交互"
	ToolsAIAgentsSettingHumanInTheLoopDesc      = "允许智能体在需要时请求用户输入"
	ToolsAIAgentsSettingUseAgents               = "启用多智能体模式"
	ToolsAIAgentsSettingUseAgentsDesc           = "允许助手调度多个专职智能体"
	ToolsAIAgentsSettingExecutionMonitor        = "启用执行监控（测试版）"
	ToolsAIAgentsSettingExecutionMonitorDesc    = "自动调用导师智能体分析执行模式"
	ToolsAIAgentsSettingSameToolLimit           = "相同工具调用阈值"
	ToolsAIAgentsSettingSameToolLimitDesc       = "触发导师复查前允许的连续相同工具调用次数"
	ToolsAIAgentsSettingTotalToolLimit          = "累计工具调用阈值"
	ToolsAIAgentsSettingTotalToolLimitDesc      = "触发导师复查前允许的累计工具调用次数"
	ToolsAIAgentsSettingMaxGeneralToolCalls     = "最大工具调用数（通用智能体）"
	ToolsAIAgentsSettingMaxGeneralToolCallsDesc = "Assistant、Primary Agent、Pentester、Coder、Installer 的最大工具调用次数"
	ToolsAIAgentsSettingMaxLimitedToolCalls     = "最大工具调用数（受限智能体）"
	ToolsAIAgentsSettingMaxLimitedToolCallsDesc = "Searcher、Enricher、Memorist 等的最大工具调用次数"
	ToolsAIAgentsSettingTaskPlanning            = "启用任务规划（测试版）"
	ToolsAIAgentsSettingTaskPlanningDesc        = "为专职智能体生成结构化执行计划"

	// help content
	ToolsAIAgentsSettingsHelp = `AI 智能体设置决定智能体如何协作、如何与用户交互以及如何进行执行控制。

基础设置：
• 启用用户交互：允许智能体在需要时请求用户输入
• 启用多智能体模式：允许助手针对复杂任务调度专职智能体

执行监控（⚠️  测试版）：
自动调用 adviser（导师）分析执行模式、检测循环、提出替代策略，避免智能体死守单一思路。阈值：连续相同调用次数（默认 5）与累计调用次数（默认 10）。

任务规划（⚠️  测试版）：
在专职智能体开始工作前生成 3-7 步执行计划，可抑制范围蔓延并提升成功率。当 adviser 采用增强配置（更强模型或最高推理档位）时效果最佳。

工具调用上限（始终生效）：
硬性上限用于防止无限循环：通用智能体默认 100，受限智能体默认 20。该限制独立于测试版功能。

参数量低于 32B 的开源模型（Qwen3.5-27B、DeepSeek-V3、Llama-3.1-70B）：
✓ 建议同时启用两项测试版功能——这对结果质量至关重要
✓ 实测结果质量相比基线提升约 2 倍
✓ 为 adviser 配置增强设置可获得最佳表现
✓ 非常适合使用本地 LLM 推理的隔离网络部署

性能影响：token 消耗与耗时增加 2-3 倍；对低于 32B 的模型质量提升约 2 倍。

⚠️  测试版提醒：功能仍在积极开发中。尽管处于测试阶段，仍推荐参数量低于 32B 的开源模型启用；若使用云端大模型 API，建议保持关闭。

注意：修改后需重启服务生效。`
)

// Search Engines screen strings
const (
	ToolsSearchEnginesFormTitle       = "搜索引擎配置"
	ToolsSearchEnginesFormDescription = "配置搜索引擎，供 AI 智能体在测试过程中收集情报"
	ToolsSearchEnginesFormName        = "搜索引擎"
	ToolsSearchEnginesFormOverview    = `可用的搜索引擎：
• DuckDuckGo —— 免费搜索引擎（无需 API 密钥）
• Sploitus —— 漏洞利用与漏洞情报库（无需 API 密钥）
• Perplexity —— 带推理能力的 AI 搜索
• Tavily —— 面向 AI 应用的搜索 API
• Traversaal —— 网页抓取与搜索
• Google Search —— 需要 API 密钥和自定义搜索引擎 ID
• Searxng —— 互联网聚合搜索引擎

API 密钥获取地址：
• Perplexity: https://www.perplexity.ai/
• Tavily: https://tavily.com/
• Traversaal: https://traversaal.ai/
• Google: https://developers.google.com/custom-search/v1/introduction`

	ToolsSearchEnginesDuckDuckGo               = "DuckDuckGo Search"
	ToolsSearchEnginesDuckDuckGoDesc           = "Enable DuckDuckGo search (no API key required)"
	ToolsSearchEnginesDuckDuckGoRegion         = "DuckDuckGo Region"
	ToolsSearchEnginesDuckDuckGoRegionDesc     = "DuckDuckGo region code (e.g., us-en, uk-en, cn-zh)"
	ToolsSearchEnginesDuckDuckGoSafeSearch     = "DuckDuckGo Safe Search"
	ToolsSearchEnginesDuckDuckGoSafeSearchDesc = "DuckDuckGo safe search (strict, moderate, off)"
	ToolsSearchEnginesDuckDuckGoTimeRange      = "DuckDuckGo Time Range"
	ToolsSearchEnginesDuckDuckGoTimeRangeDesc  = "DuckDuckGo time range (d: day, w: week, m: month, y: year)"
	ToolsSearchEnginesSploitus                 = "Sploitus Search"
	ToolsSearchEnginesSploitusDesc             = "Enable Sploitus search for exploits and vulnerabilities (no API key required)"
	ToolsSearchEnginesPerplexityKey            = "Perplexity API Key"
	ToolsSearchEnginesPerplexityKeyDesc        = "API key for Perplexity AI search"
	ToolsSearchEnginesTavilyKey                = "Tavily API Key"
	ToolsSearchEnginesTavilyKeyDesc            = "API key for Tavily search service"
	ToolsSearchEnginesTraversaalKey            = "Traversaal API Key"
	ToolsSearchEnginesTraversaalKeyDesc        = "API key for Traversaal web scraping"
	ToolsSearchEnginesGoogleKey                = "Google Search API Key"
	ToolsSearchEnginesGoogleKeyDesc            = "Google Custom Search API key"
	ToolsSearchEnginesGoogleCX                 = "Google Search Engine ID"
	ToolsSearchEnginesGoogleCXDesc             = "Google Custom Search Engine ID"
	ToolsSearchEnginesGoogleLR                 = "Google Language Restriction"
	ToolsSearchEnginesGoogleLRDesc             = "Google Search Engine language restriction (e.g., lang_en, lang_cn, etc.)"
	ToolsSearchEnginesSearxngURL               = "Searxng Search URL"
	ToolsSearchEnginesSearxngURLDesc           = "Searxng search engine URL"
	ToolsSearchEnginesSearxngCategories        = "Searxng Search Categories"
	ToolsSearchEnginesSearxngCategoriesDesc    = "Searxng search engine categories (e.g., general, it, web, news, technology, science, health, other)"
	ToolsSearchEnginesSearxngLanguage          = "Searxng Search Language"
	ToolsSearchEnginesSearxngLanguageDesc      = "Searxng search engine language (en, ch, fr, de, it, es, pt, ru, zh, empty for all languages)"
	ToolsSearchEnginesSearxngSafeSearch        = "Searxng Safe Search"
	ToolsSearchEnginesSearxngSafeSearchDesc    = "Searxng search engine safe search (0: off, 1: moderate, 2: strict)"
	ToolsSearchEnginesSearxngTimeRange         = "Searxng Time Range"
	ToolsSearchEnginesSearxngTimeRangeDesc     = "Searxng search engine time range (day, month, year)"
	ToolsSearchEnginesSearxngTimeout           = "Searxng Timeout"
	ToolsSearchEnginesSearxngTimeoutDesc       = "Searxng request timeout in seconds"
)

// Scraper screen strings
const (
	ToolsScraperFormTitle       = "Scraper Configuration"
	ToolsScraperFormDescription = "Configure web scraping service"
	ToolsScraperFormName        = "Scraper"
	ToolsScraperFormOverview    = `Web scraper service for content extraction and analysis using vxcontrol/scraper Docker image.

Modes:
• Embedded - Run local scraper container (recommended)
• External - Use external scraper services
• Disabled - No web scraping capabilities

Docker image: https://hub.docker.com/r/vxcontrol/scraper

The scraper supports:
• Public URL access for external links
• Private URL access for internal/local links
• Content extraction and analysis
• Multiple output formats`

	ToolsScraperModeTitle                 = "Scraper Mode"
	ToolsScraperModeDesc                  = "Select how the scraper service should operate"
	ToolsScraperEmbedded                  = "Embedded Container"
	ToolsScraperExternal                  = "External Service"
	ToolsScraperDisabled                  = "Disabled"
	ToolsScraperPublicURL                 = "Public Scraper URL"
	ToolsScraperPublicURLDesc             = "URL for scraping public/external websites. If empty, the same value as private URL will be used."
	ToolsScraperPublicURLEmbeddedDesc     = "URL for embedded scraper (optional override). If empty, the same value as private URL will be used."
	ToolsScraperPrivateURL                = "Private Scraper URL"
	ToolsScraperPrivateURLDesc            = "URL for scraping private/internal websites"
	ToolsScraperPublicUsername            = "Public URL Username"
	ToolsScraperPublicUsernameDesc        = "Username for public scraper access"
	ToolsScraperPublicPassword            = "Public URL Password"
	ToolsScraperPublicPasswordDesc        = "Password for public scraper access"
	ToolsScraperPrivateUsername           = "Private URL Username"
	ToolsScraperPrivateUsernameDesc       = "Username for private scraper access"
	ToolsScraperPrivatePassword           = "Private URL Password"
	ToolsScraperPrivatePasswordDesc       = "Password for private scraper access"
	ToolsScraperLocalUsername             = "Local URL Username"
	ToolsScraperLocalUsernameDesc         = "Username for embedded scraper service"
	ToolsScraperLocalPassword             = "Local URL Password"
	ToolsScraperLocalPasswordDesc         = "Password for embedded scraper service"
	ToolsScraperMaxConcurrentSessions     = "Max Concurrent Sessions"
	ToolsScraperMaxConcurrentSessionsDesc = "Maximum number of concurrent scraping sessions"
	ToolsScraperEmbeddedHelp              = "Embedded mode runs a local scraper container that can access both public and private resources. The default configuration uses https://someuser:somepass@scraper/."
	ToolsScraperExternalHelp              = "External mode uses separate scraper services. Configure different URLs for public and private access as needed."
	ToolsScraperDisabledHelp              = "Scraper is disabled. Web content extraction and analysis capabilities will not be available."
)

// Docker Environment screen strings
const (
	ToolsDockerFormTitle       = "Docker Environment Configuration"
	ToolsDockerFormDescription = "Configure Docker environment for worker containers"
	ToolsDockerFormName        = "Docker Environment"
	ToolsDockerFormOverview    = `• Worker Isolation - Containers provide security boundaries for tasks
• Network Capabilities - Enable privileged network operations for pentesting
• Container Management - Control how workers access Docker daemon
• Storage Configuration - Define workspace and artifact storage
• Image Selection - Set default images for different task types

Critical for penetration testing workflows requiring network scanning, custom tools, and secure task isolation.`

	// General help text
	ToolsDockerGeneralHelp = `Each AI agent task runs in an isolated Docker container with two ports (28000-32000 range) automatically allocated per flow. Worker containers are created on-demand from default images or agent-selected ones.

Basic setup requires enabling capabilities: Docker Access allows spawning additional containers for specialized tools, while Network Admin grants low-level network permissions essential for scanning tools like nmap.

Storage operates via Docker volumes by default, or host directories when Work Directory is specified. Connection settings control the Docker daemon location - local socket for standard setups, or remote TCP with TLS for distributed environments.

Default images serve as fallbacks: general tasks use standard images, while security testing defaults to pentesting-focused containers. Public IP enables reverse shell attacks by providing workers with a reachable address for target callbacks. Usually it's a local interface address of the host machine with Docker daemon running for the workers containers.

Configuration combines based on scenario: enable both capabilities for full pentesting, use Work Directory for persistent artifacts, or configure remote connection for isolated Docker environments.`

	// Container capabilities
	ToolsDockerInside       = "Docker Access"
	ToolsDockerInsideDesc   = "Allow workers to manage Docker containers"
	ToolsDockerNetAdmin     = "Network Admin"
	ToolsDockerNetAdminDesc = "Grant NET_ADMIN capability for network scanning tools like nmap"

	// Connection settings
	ToolsDockerSocket       = "Docker Socket"
	ToolsDockerSocketDesc   = "Path to Docker socket on host filesystem"
	ToolsDockerNetwork      = "Docker Network"
	ToolsDockerNetworkDesc  = "Custom network name for worker containers, or 'host' for direct host network access"
	ToolsDockerPublicIP     = "Public IP Address"
	ToolsDockerPublicIPDesc = "Public IP for reverse connections in OOB attacks"

	// Storage configuration
	ToolsDockerWorkDir     = "Work Directory"
	ToolsDockerWorkDirDesc = "Host directory for worker filesystems (default: Docker volumes)"

	// Default images
	ToolsDockerDefaultImage               = "Default Image"
	ToolsDockerDefaultImageDesc           = "Default Docker image for general tasks"
	ToolsDockerDefaultImageForPentest     = "Pentesting Image"
	ToolsDockerDefaultImageForPentestDesc = "Default Docker image for security testing tasks"

	// TLS connection settings (optional)
	ToolsDockerHost          = "Docker Host"
	ToolsDockerHostDesc      = "Docker daemon connection (unix:// or tcp://)"
	ToolsDockerTLSVerify     = "TLS Verification"
	ToolsDockerTLSVerifyDesc = "Enable TLS verification for Docker connection"
	ToolsDockerCertPath      = "TLS Certificates"
	ToolsDockerCertPathDesc  = "Directory containing ca.pem, cert.pem, key.pem files"

	// Help content for specific configurations
	ToolsDockerInsideHelp = `Docker Access enables workers to spawn additional containers for specialized tools and environments. Required when tasks need custom software not available in default images.

When enabled, workers can pull and run any Docker image, providing maximum flexibility for complex testing scenarios.`

	ToolsDockerNetAdminHelp = `Network Admin capability allows workers to perform low-level network operations essential for penetration testing.

Required for:
• Network scanning with nmap, masscan
• Custom packet crafting
• Network interface manipulation
• Raw socket operations

Critical for comprehensive security assessments.`

	ToolsDockerSocketHelp = `Docker Socket path defines how workers access the Docker daemon. Use only file path to the socket file. Used with Docker Access to enable container management.

For enhanced security, consider using docker-in-docker (DinD) instead of exposing the main Docker daemon directly to workers.
When using DinD, use the path to the Docker socket file of the DinD container which binded to the host filesystem.

Example: /var/run/docker.sock`

	ToolsDockerNetworkHelp = `Docker Network controls network isolation mode for worker containers:

Bridge Mode (custom network name):
• Isolated communication between containers
• Port forwarding from container to host
• Enhanced security boundaries
• Network-based monitoring and filtering
• Recommended for most use cases

主机模式（值为 'host'）：
• 直接访问主机网络接口
• 无端口转发，端口直接绑定到主机
• 原始数据包操作所必需
• 支持高级网络测试能力
• 隔离性较低，请谨慎使用

示例：
• 'pentagi-network' —— 创建隔离的桥接网络
• 'host' —— 启用直接的主机网络访问

安全提示：主机网络模式会降低容器隔离性。仅在高级渗透测试任务确实需要直接访问网络栈时使用。`

	ToolsDockerPublicIPHelp = `公网 IP 地址用于启用带外（OOB）攻击技术，为工作容器提供可供反向连接的可达地址。

工作容器会自动获得两个映射到该 IP 的随机端口（28000-32000 范围），用于接收来自被利用目标的回连。

默认情况下，智能体会尝试通过 api.ipify.org、ipinfo.io/ip 或 ifconfig.me 服务获取公网地址。`

	ToolsDockerWorkDirHelp = `工作目录指定工作容器存储所在的主机文件系统位置。设置后将以主机目录挂载替代默认的 Docker 卷。

优势：
• 重启后数据仍然保留
• 可直接访问文件系统
• 更便于管理产物文件
• 可自定义备份策略

默认情况下为每个工作容器使用独立的 Docker 卷。

示例：/path/to/workdir/`

	ToolsDockerDefaultImageHelp = `当任务需求未指定特定容器镜像时，默认镜像作为工作容器的回退选项。

应包含通用任务所需的基础工具与实用程序。默认值：debian:latest`

	ToolsDockerDefaultImageForPentestHelp = `渗透测试镜像作为安全测试任务的默认镜像，应包含完备的安全工具与实用程序。

推荐使用 Kali Linux、Parrot Security 或自定义的安全专用容器。默认值：vxcontrol/kali-linux`

	ToolsDockerHostHelp = `Docker Host 用于启动主工作容器，并覆盖默认的 Docker 守护进程连接。支持 Unix 套接字和 TCP 连接。

示例：
• unix:///var/run/docker.sock（本地）
• tcp://docker-host:2376（远程）

远程连接请启用 TLS。`

	ToolsDockerTLSVerifyHelp = `TLS 验证用于保护通过 TCP 建立的 Docker 守护进程连接。强烈建议对远程 Docker 主机启用。

需要在指定的证书目录中提供有效证书。`

	ToolsDockerCertPathHelp = `TLS 证书目录必须包含：
• ca.pem —— 证书颁发机构证书
• cert.pem —— 客户端证书
• key.pem —— 私钥

使用 TLS 管理工作容器时，安全的远程 Docker 连接需要这些文件。

示例：/path/to/certs`
)

// Embedder form strings
const (
	EmbedderFormTitle       = "嵌入模型配置"
	EmbedderFormDescription = "配置用于语义搜索和知识存储的文本向量化"
	EmbedderFormName        = "嵌入模型"
	EmbedderFormOverview    = `文本嵌入将文档转换为向量，用于语义搜索和知识存储。
不同提供商提供的模型在能力和价格上各有差异。

请谨慎选择：更换提供商需要对所有已存储数据重新建立索引。`

	EmbedderFormProvider     = "嵌入提供商"
	EmbedderFormProviderDesc = "选择用于文本向量化的提供商。嵌入用于语义搜索和知识存储。"

	EmbedderFormURL     = "API 端点地址"
	EmbedderFormURLDesc = "自定义 API 端点（留空则使用默认值）"

	EmbedderFormAPIKey     = "API 密钥"
	EmbedderFormAPIKeyDesc = "提供商的身份验证密钥（Ollama 无需填写）"

	EmbedderFormModel     = "模型名称"
	EmbedderFormModelDesc = "要使用的具体嵌入模型（留空则使用提供商默认值）"

	EmbedderFormBatchSize     = "批处理大小"
	EmbedderFormBatchSizeDesc = "单批处理的文档数量（1-1000）"

	EmbedderFormStripNewLines     = "去除换行符"
	EmbedderFormStripNewLinesDesc = "嵌入前移除文本中的换行符（true/false）"

	EmbedderFormMaxTextBytes     = "文本字节上限"
	EmbedderFormMaxTextBytesDesc = "发送到嵌入 API 的每个文本块的最大字节数（例如 8192）"

	EmbedderFormHelpTitle   = "嵌入配置"
	EmbedderFormHelpContent = `配置用于语义搜索和知识存储的文本向量化。

若未配置具体的嵌入设置，系统将使用 OpenAI 嵌入，并采用 LLM 提供商中配置的 API 密钥。

更换提供商需谨慎：不同嵌入模型生成的向量互不兼容，需要重建数据库索引。`

	EmbedderFormHelpOpenAI      = "OpenAI：最可靠的选择，质量优异。若此处未填写，将使用 LLM 提供商中的 API 密钥。"
	EmbedderFormHelpOllama      = "Ollama：本地嵌入，无需 API 密钥。需要运行 Ollama 服务器。"
	EmbedderFormHelpHuggingFace = "HuggingFace：开源模型，需要 API 密钥。"
	EmbedderFormHelpGoogleAI    = "Google AI：高质量嵌入，需要 API 密钥。"

	// Provider names and descriptions
	EmbedderProviderDefault         = "默认（OpenAI）"
	EmbedderProviderDefaultDesc     = "使用 OpenAI 嵌入，并采用 LLM 提供商配置中的 API 密钥"
	EmbedderProviderOpenAI          = "OpenAI"
	EmbedderProviderOpenAIDesc      = "OpenAI 文本嵌入 API（text-embedding-3-small、ada-002）"
	EmbedderProviderOllama          = "Ollama"
	EmbedderProviderOllamaDesc      = "本地 Ollama 服务器，运行开源嵌入模型"
	EmbedderProviderMistral         = "Mistral"
	EmbedderProviderMistralDesc     = "Mistral AI 嵌入模型"
	EmbedderProviderJina            = "Jina"
	EmbedderProviderJinaDesc        = "Jina AI 嵌入 API"
	EmbedderProviderHuggingFace     = "HuggingFace"
	EmbedderProviderHuggingFaceDesc = "HuggingFace 推理 API，用于嵌入模型"
	EmbedderProviderGoogleAI        = "Google AI"
	EmbedderProviderGoogleAIDesc    = "Google AI 嵌入模型（embedding-001）"
	EmbedderProviderVoyageAI        = "VoyageAI"
	EmbedderProviderVoyageAIDesc    = "VoyageAI 嵌入 API"
	EmbedderProviderDisabled        = "已禁用"
	EmbedderProviderDisabledDesc    = "完全禁用嵌入功能"

	// Provider-specific placeholders and help
	EmbedderURLPlaceholderOpenAI      = "https://api.openai.com/v1"
	EmbedderURLPlaceholderOllama      = "http://localhost:11434"
	EmbedderURLPlaceholderMistral     = "https://api.mistral.ai/v1"
	EmbedderURLPlaceholderJina        = "https://api.jina.ai/v1"
	EmbedderURLPlaceholderHuggingFace = "https://api-inference.huggingface.co"
	EmbedderURLPlaceholderGoogleAI    = "Not supported - uses default endpoint"
	EmbedderURLPlaceholderVoyageAI    = "Not supported - uses default endpoint"

	EmbedderAPIKeyPlaceholderOllama      = "Not required for local models"
	EmbedderAPIKeyPlaceholderMistral     = "Mistral API key"
	EmbedderAPIKeyPlaceholderJina        = "Jina API key"
	EmbedderAPIKeyPlaceholderHuggingFace = "HuggingFace API key"
	EmbedderAPIKeyPlaceholderGoogleAI    = "Google AI API key"
	EmbedderAPIKeyPlaceholderVoyageAI    = "VoyageAI API key"
	EmbedderAPIKeyPlaceholderDefault     = "API key for the provider"

	EmbedderModelPlaceholderOpenAI      = "text-embedding-3-small"
	EmbedderModelPlaceholderOllama      = "nomic-embed-text"
	EmbedderModelPlaceholderMistral     = "mistral-embed"
	EmbedderModelPlaceholderJina        = "jina-embeddings-v2-base-en"
	EmbedderModelPlaceholderHuggingFace = "sentence-transformers/all-MiniLM-L6-v2"
	EmbedderModelPlaceholderGoogleAI    = "gemini-embedding-001"
	EmbedderModelPlaceholderVoyageAI    = "voyage-2"
	EmbedderModelPlaceholderDefault     = "Model name"

	// Provider IDs for internal use
	EmbedderProviderIDDefault     = "default"
	EmbedderProviderIDOpenAI      = "openai"
	EmbedderProviderIDOllama      = "ollama"
	EmbedderProviderIDMistral     = "mistral"
	EmbedderProviderIDJina        = "jina"
	EmbedderProviderIDHuggingFace = "huggingface"
	EmbedderProviderIDGoogleAI    = "googleai"
	EmbedderProviderIDVoyageAI    = "voyageai"
	EmbedderProviderIDDisabled    = "none"

	EmbedderHelpGeneral = `Embeddings convert text into vectors for semantic search and knowledge storage. This enables PentAGI to understand meaning rather than just keywords, making search results more relevant and intelligent.

Key benefits:
• Find documents by meaning, not exact words
• Build a smart knowledge base from pentesting results
• Enable AI agents to locate relevant information quickly
• Support advanced reasoning with contextual data

Choose Ollama for completely local processing - your data never leaves your infrastructure. Other providers offer cloud-based processing with different model capabilities and pricing.

Configure carefully as changing providers requires rebuilding the entire knowledge base.`

	EmbedderHelpAttentionPrefix = "Important:"
	EmbedderHelpAttention       = `Different embedding providers create incompatible vectors. Changing providers or models will break existing semantic search.

You must flush or reindex your entire knowledge base using the etester utility:
• Run 'etester flush' to clear old embeddings
• Run 'etester reindex' to rebuild with new provider
• This process can take significant time for large datasets`

	EmbedderHelpAttentionSuffix = `Only change providers if absolutely necessary.`

	// Provider help texts
	EmbedderHelpDefault = `Default mode uses OpenAI embeddings with the API key configured in LLM Providers.

This is the recommended option for most users as it requires no additional configuration if you already have OpenAI set up.`

	EmbedderHelpOpenAI = `Direct OpenAI API access for embedding generation.

Get your API key from:
https://platform.openai.com/api-keys

Recommended models:
• text-embedding-3-small (cost-effective, 1536 dimensions)
• text-embedding-3-large (highest quality, 3072 dimensions)
• text-embedding-ada-002 (legacy, still supported)`

	EmbedderHelpOllama = `Local Ollama server for open-source embedding models.

Popular embedding models:
• nomic-embed-text (recommended, 768 dimensions)
• mxbai-embed-large (large model, 1024 dimensions)
• snowflake-arctic-embed (multilingual support)

Install Ollama from:
https://ollama.com/

Start with: ollama pull nomic-embed-text`

	EmbedderHelpMistral = `Mistral AI embedding models via API.

Get your API key from:
https://console.mistral.ai/

Uses Mistral's embedding model with fixed configuration.
No model selection required - uses the default embedding model.`

	EmbedderHelpJina = `Jina AI embedding API with specialized models.

Get your API key from:
https://jina.ai/

Recommended models:
• jina-embeddings-v2-base-en (general purpose, 768 dimensions)
• jina-embeddings-v2-small-en (lightweight, 512 dimensions)
• jina-embeddings-v2-base-code (code-specific embeddings)`

	EmbedderHelpHuggingFace = `HuggingFace Inference API for open-source models.

Get your API key from:
https://huggingface.co/settings/tokens

Popular models:
• sentence-transformers/all-MiniLM-L6-v2 (384 dimensions)
• sentence-transformers/all-mpnet-base-v2 (768 dimensions)
• intfloat/e5-large-v2 (1024 dimensions)`

	EmbedderHelpGoogleAI = `Google AI embedding models (Gemini).

Get your API key from:
https://aistudio.google.com/app/apikey

Available models:
• gemini-embedding-001 (latest model, 768 dimensions)
• text-embedding-004 (legacy Vertex AI model)

Uses Google's fixed endpoint - URL configuration not supported.`

	EmbedderHelpVoyageAI = `VoyageAI embedding API optimized for retrieval.

Get your API key from:
https://www.voyageai.com/

Recommended models:
• voyage-2 (general purpose, 1024 dimensions)
• voyage-large-2 (highest quality, 1536 dimensions)
• voyage-code-2 (code embeddings, 1536 dimensions)`

	EmbedderHelpDisabled = `Disables all embedding functionality.

This will:
• Disable semantic search capabilities
• Turn off knowledge storage vectorization
• Reduce memory and computational requirements

Only recommended if embeddings are not needed for your use case.`
)

// Development and Mock Screen constants
const (
	MockScreenTitle       = "开发中界面"
	MockScreenDescription = "此界面尚在开发中"
)

// Apply Changes screen constants
const (
	ApplyChangesFormTitle       = "应用配置更改"
	ApplyChangesFormName        = "应用更改"
	ApplyChangesFormDescription = "检查并应用你的配置更改"

	// Apply Changes overview and help
	ApplyChangesFormOverview = `在此界面可以查看所有待应用的配置更改，并将它们应用到你的 PentAGI 安装中。

应用更改时，系统将：
• 把所有修改过的环境变量保存到 .env 文件
• 使用新配置重启受影响的服务
• 按需安装额外组件`

	// Apply Changes status messages
	ApplyChangesNotStarted     = "配置更改已准备就绪，可以应用"
	ApplyChangesInProgress     = "正在应用配置更改...\n"
	ApplyChangesCompleted      = "配置更改已成功应用\n"
	ApplyChangesFailed         = "配置更改执行失败"
	ApplyChangesResetCompleted = "配置更改已成功重置\n"

	ApplyChangesTerminalIsNotInitialized = "终端尚未初始化"

	// Apply Changes instructions
	ApplyChangesInstructions = `按 Enter 开始应用配置更改。`

	ApplyChangesNoChanges = "没有待应用的配置更改"

	// Apply Changes installation status
	ApplyChangesInstallNotFound = `本系统当前尚未安装 PentAGI。

将执行以下操作：
• 设置并校验 Docker 环境
• 创建 docker-compose.yml 文件
• 安装并启动 PentAGI 核心服务`

	ApplyChangesInstallFoundLangfuse      = `• 安装 Langfuse 可观测性组件栈（docker-compose-langfuse.yml）`
	ApplyChangesInstallFoundObservability = `• 安装包含 Grafana、VictoriaMetrics 和 Jaeger 的完整可观测性组件栈（docker-compose-observability.yml）`

	ApplyChangesUpdateFound = `本系统当前已安装 PentAGI。

将执行以下操作：
• 更新 .env 文件中的环境变量
• 重建并重启受影响的 Docker 容器
• 将新配置应用到正在运行的服务`

	// Apply Changes warnings and notes
	ApplyChangesWarningCritical = "⚠️  检测到关键更改 —— 服务将被重启"
	ApplyChangesWarningSecrets  = "🔒 检测到敏感信息 —— 将以安全方式存储"
	ApplyChangesNoteBackup      = "💾 更改前会先备份当前配置"
	ApplyChangesNoteTime        = "⏱️  视所选组件而定，此过程通常不超过一分钟"

	// Apply Changes progress messages
	ApplyChangesStageValidation = "正在校验环境与依赖..."
	ApplyChangesStageBackup     = "正在创建配置备份..."
	ApplyChangesStageEnvFile    = "正在更新环境配置文件..."
	ApplyChangesStageCompose    = "正在生成 Docker Compose 文件..."
	ApplyChangesStageDocker     = "正在管理 Docker 容器..."
	ApplyChangesStageServices   = "正在启动服务..."
	ApplyChangesStageComplete   = "配置更改已成功应用"

	// Apply Changes change list headers
	ApplyChangesChangesTitle  = "待应用的配置更改"
	ApplyChangesChangesCount  = "更改总数：%d"
	ApplyChangesChangesMasked = "（出于安全已隐藏）"
	ApplyChangesChangesEmpty  = "没有需要应用的更改"

	// Apply Changes help content
	ApplyChangesHelpTitle   = "应用配置更改"
	ApplyChangesHelpContent = `应用更改前请务必确认当前配置。`
)

// apply changes integrity prompt
const (
	ApplyChangesIntegrityPromptTitle   = "文件完整性检查"
	ApplyChangesIntegrityPromptMessage = "检测到过期文件。\n是否将它们更新到最新版本？"
	ApplyChangesIntegrityOutdatedList  = "过期文件：\n%s\n确认更新？(y/n)"
	ApplyChangesIntegrityChecking      = "正在收集文件完整性信息..."
	ApplyChangesIntegrityNoOutdated    = "未发现过期文件，继续应用更改。"
)

// Maintenance Screen constants
const (
	MaintenanceTitle       = "System Maintenance"
	MaintenanceDescription = "Manage PentAGI services and perform maintenance operations"
	MaintenanceName        = "Maintenance"
	MaintenanceOverview    = `Perform system maintenance operations for PentAGI.

Available operations depend on the current system state and will only be shown when applicable.

Operations include:
• Service lifecycle management (Start/Stop/Restart)
• Component updates and downloads
• System reset and cleanup
• Container and image management

Each operation will provide real-time status updates and confirmation when required.`

	// Maintenance menu items
	MaintenanceStartPentagi            = "Start PentAGI"
	MaintenanceStartPentagiDesc        = "Start all configured PentAGI services"
	MaintenanceStopPentagi             = "Stop PentAGI"
	MaintenanceStopPentagiDesc         = "Stop all running PentAGI services"
	MaintenanceRestartPentagi          = "Restart PentAGI"
	MaintenanceRestartPentagiDesc      = "Restart all PentAGI services"
	MaintenanceDownloadWorkerImage     = "Download Worker Image"
	MaintenanceDownloadWorkerImageDesc = "Download pentesting container image for worker tasks"
	MaintenanceUpdateWorkerImage       = "Update Worker Image"
	MaintenanceUpdateWorkerImageDesc   = "Update pentesting container image to latest version"
	MaintenanceUpdatePentagi           = "Update PentAGI"
	MaintenanceUpdatePentagiDesc       = "Update PentAGI to the latest version"
	MaintenanceUpdateInstaller         = "Update Installer"
	MaintenanceUpdateInstallerDesc     = "Update this installer to the latest version"
	MaintenanceFactoryReset            = "Factory Reset"
	MaintenanceFactoryResetDesc        = "Reset PentAGI to factory defaults"
	MaintenanceRemovePentagi           = "Remove PentAGI"
	MaintenanceRemovePentagiDesc       = "Remove PentAGI containers but keep data"
	MaintenancePurgePentagi            = "Purge PentAGI"
	MaintenancePurgePentagiDesc        = "Completely remove PentAGI including all data"
	MaintenanceResetPassword           = "Reset Admin Password"
	MaintenanceResetPasswordDesc       = "Reset the administrator password for PentAGI"
)

// Reset Password Screen constants
const (
	ResetPasswordFormTitle       = "Reset Admin Password"
	ResetPasswordFormDescription = "Reset the administrator password for PentAGI"
	ResetPasswordFormName        = "Reset Password"
	ResetPasswordFormOverview    = `Reset the password for the default administrator account (admin@pentagi.com).

This operation requires PentAGI to be running and will update the password in the PostgreSQL database.

Enter your new password twice to confirm and press Enter to apply the change.

Password requirements:
• Minimum 5 characters
• Both password fields must match`

	// Form fields
	ResetPasswordNewPassword         = "New Password"
	ResetPasswordNewPasswordDesc     = "Enter the new administrator password"
	ResetPasswordConfirmPassword     = "Confirm Password"
	ResetPasswordConfirmPasswordDesc = "Re-enter the new password to confirm"

	// Status messages
	ResetPasswordNotAvailable = "PentAGI must be running to reset password"
	ResetPasswordAvailable    = "Password reset is available"
	ResetPasswordInProgress   = "Resetting password..."
	ResetPasswordSuccess      = "Password has been successfully reset"
	ResetPasswordErrorPrefix  = "Error: "

	// Validation errors
	ResetPasswordErrorEmptyPassword = "Password cannot be empty"
	ResetPasswordErrorShortPassword = "Password must be at least 5 characters long"
	ResetPasswordErrorMismatch      = "Passwords do not match"

	// Help content
	ResetPasswordHelpContent = `Reset the administrator password for accessing PentAGI.

This operation:
• Updates the password for admin@pentagi.com account
• Sets the user status to 'active'
• Requires PentAGI database to be accessible
• Does not affect other user accounts

The password change takes effect immediately after successful completion.

Enter the same password in both fields and press Enter to confirm the change.`
)

// Processor Operation Form constants
const (
	// Dynamic title templates
	ProcessorOperationFormTitle       = "%s"
	ProcessorOperationFormDescription = "Execute %s operation"
	ProcessorOperationFormName        = "%s"

	// Common status messages
	ProcessorOperationNotStarted = "Ready to execute %s operation"
	ProcessorOperationInProgress = "Executing %s operation...\n"
	ProcessorOperationCompleted  = "%s operation completed successfully\n"
	ProcessorOperationFailed     = "Failed to execute %s operation"

	// Confirmation messages
	ProcessorOperationConfirmation = "Are you sure you want to %s?"
	ProcessorOperationPressEnter   = "Press Enter to %s"
	ProcessorOperationPressYN      = "Press Y to confirm, N to cancel"
	// Short notice without hotkeys (for static help panel)
	ProcessorOperationRequiresConfirmationShort = "This operation requires confirmation"
	// Additional terminal messages
	ProcessorOperationCancelled = "Operation cancelled"
	ProcessorOperationUnknown   = "Unknown operation: %s"

	// Operation specific messages
	ProcessorOperationStarting    = "Starting services..."
	ProcessorOperationStopping    = "Stopping services..."
	ProcessorOperationRestarting  = "Restarting services..."
	ProcessorOperationDownloading = "Downloading images..."
	ProcessorOperationUpdating    = "Updating components..."
	ProcessorOperationResetting   = "Resetting to factory defaults..."
	ProcessorOperationRemoving    = "Removing containers..."
	ProcessorOperationPurging     = "Purging all data..."
	ProcessorOperationInstalling  = "Installing PentAGI services..."

	// Help text templates
	ProcessorOperationHelpTitle           = "%s Operation"
	ProcessorOperationHelpContent         = "This operation will %s."
	ProcessorOperationHelpContentDownload = "This operation will download %s components."
	ProcessorOperationHelpContentUpdate   = "This operation will update %s components."
	// Generic title/description/builders for dynamic operations
	OperationTitleInstallPentagi    = "Install PentAGI"
	OperationDescInstallPentagi     = "Install and configure PentAGI services"
	OperationTitleDownload          = "Download %s"
	OperationDescDownloadComponents = "Download %s components"
	OperationTitleUpdate            = "Update %s"
	OperationDescUpdateToLatest     = "Update %s to latest version"
	OperationTitleExecute           = "Execute %s"
	OperationDescExecuteOn          = "Execute %s on %s"
	OperationProgressExecuting      = "Executing %s..."

	// Terminal not initialized
	ProcessorOperationTerminalNotInitialized = "Terminal is not initialized"
)

// Operation-specific help texts
const (
	ProcessorHelpInstallPentagi = `This will:
• Deploy Docker containers for selected services
• Configure networking and volumes
• Start all enabled services
• Set up monitoring if configured

Installation will use your current configuration settings.`

	ProcessorHelpStartPentagi = `This will:
• Core PentAGI API and web interface
• Configured Langfuse analytics (if enabled)
• Observability stack (if enabled)

Services will be started in the correct dependency order.`

	ProcessorHelpStopPentagi = `This will:
• Gracefully shutdown containers
• Preserve all data and configurations
• Network connections will be closed

You can restart services later without losing any data.`

	ProcessorHelpRestartPentagi = `This will:
• Stop running containers
• Apply any configuration changes
• Start services with fresh state

Useful after configuration updates or to resolve issues.`

	ProcessorHelpDownloadWorkerImage = `This large image (6GB+) contains:
• Kali Linux tools and utilities
• Security testing frameworks
• Network analysis software

Required for pentesting operations.`

	ProcessorHelpUpdateWorkerImage = `This will:
• Pull the latest pentesting image
• Update security tools and frameworks
• Preserve existing worker containers

Note: This is a large download (6GB+).`

	ProcessorHelpUpdatePentagi = `This will:
• Download latest container images
• Perform rolling update of services
• Preserve all data and configurations

Services will be briefly unavailable during update.`

	ProcessorHelpUpdateInstaller = `This will:
• Download the latest installer binary
• Replace the current installer
• Exit for manual restart

You'll need to restart the installer after update.`

	ProcessorHelpFactoryReset = `⚠️  WARNING: This operation will:
• Remove all containers and networks
• Delete all configuration files
• Clear stored data and volumes
• Restore default settings

This action cannot be undone!`

	ProcessorHelpRemovePentagi = `This will:
• Stop and remove all containers
• Remove Docker networks
• Preserve volumes and data
• Keep configuration files

You can reinstall later without losing data.`

	ProcessorHelpPurgePentagi = `⚠️  WARNING: This will permanently delete:
• All containers and images
• All data volumes
• All configuration files
• All stored results

This action cannot be undone!`
)

// environment variable descriptions (centralized)
const (
	EnvDesc_OPEN_AI_KEY                       = "OpenAI API Key"
	EnvDesc_OPEN_AI_SERVER_URL                = "OpenAI Server URL"
	EnvDesc_ANTHROPIC_API_KEY                 = "Anthropic API Key"
	EnvDesc_ANTHROPIC_SERVER_URL              = "Anthropic Server URL"
	EnvDesc_GEMINI_API_KEY                    = "Google Gemini API Key"
	EnvDesc_GEMINI_SERVER_URL                 = "Gemini Server URL"
	EnvDesc_BEDROCK_DEFAULT_AUTH              = "AWS Bedrock Use Default Credential Chain"
	EnvDesc_BEDROCK_BEARER_TOKEN              = "AWS Bedrock Bearer Token"
	EnvDesc_BEDROCK_ACCESS_KEY_ID             = "AWS Bedrock Access Key ID"
	EnvDesc_BEDROCK_SECRET_ACCESS_KEY         = "AWS Bedrock Secret Access Key"
	EnvDesc_BEDROCK_SESSION_TOKEN             = "AWS Bedrock Session Token"
	EnvDesc_BEDROCK_REGION                    = "AWS Bedrock Region"
	EnvDesc_BEDROCK_SERVER_URL                = "AWS Bedrock Custom Endpoint URL"
	EnvDesc_OLLAMA_SERVER_URL                 = "Ollama Server URL"
	EnvDesc_OLLAMA_SERVER_API_KEY             = "Ollama Server API Key (Cloud)"
	EnvDesc_OLLAMA_SERVER_MODEL               = "Ollama Default Model"
	EnvDesc_OLLAMA_SERVER_CONFIG_PATH         = "Ollama Container Config Path"
	EnvDesc_OLLAMA_SERVER_PULL_MODELS_TIMEOUT = "Ollama Model Pull Timeout"
	EnvDesc_OLLAMA_SERVER_PULL_MODELS_ENABLED = "Ollama Auto-pull Models"
	EnvDesc_OLLAMA_SERVER_LOAD_MODELS_ENABLED = "Ollama Load Models List"
	EnvDesc_DEEPSEEK_API_KEY                  = "DeepSeek API Key"
	EnvDesc_DEEPSEEK_SERVER_URL               = "DeepSeek Server URL"
	EnvDesc_DEEPSEEK_PROVIDER                 = "DeepSeek Provider Name Prefix (for LiteLLM, e.g., 'deepseek')"
	EnvDesc_GLM_API_KEY                       = "GLM API Key"
	EnvDesc_GLM_SERVER_URL                    = "GLM Server URL"
	EnvDesc_GLM_PROVIDER                      = "GLM Provider Name Prefix (for LiteLLM, e.g., 'zai')"
	EnvDesc_KIMI_API_KEY                      = "Kimi API Key"
	EnvDesc_KIMI_SERVER_URL                   = "Kimi Server URL"
	EnvDesc_KIMI_PROVIDER                     = "Kimi Provider Name Prefix (for LiteLLM, e.g., 'moonshot')"
	EnvDesc_QWEN_API_KEY                      = "Qwen API Key"
	EnvDesc_QWEN_SERVER_URL                   = "Qwen Server URL"
	EnvDesc_QWEN_PROVIDER                     = "Qwen Provider Name Prefix (for LiteLLM, e.g., 'dashscope')"
	EnvDesc_LLM_SERVER_URL                    = "Custom LLM Server URL"
	EnvDesc_LLM_SERVER_KEY                    = "Custom LLM API Key"
	EnvDesc_LLM_SERVER_MODEL                  = "Custom LLM Model"
	EnvDesc_LLM_SERVER_CONFIG_PATH            = "Custom LLM Container Config Path"
	EnvDesc_LLM_SERVER_LEGACY_REASONING       = "Custom LLM Legacy Reasoning"
	EnvDesc_LLM_SERVER_PRESERVE_REASONING     = "Custom LLM Preserve Reasoning Content"
	EnvDesc_LLM_SERVER_PROVIDER               = "Custom LLM Provider Name"

	EnvDesc_LANGFUSE_LISTEN_IP   = "Langfuse Listen IP"
	EnvDesc_LANGFUSE_LISTEN_PORT = "Langfuse Listen Port"
	EnvDesc_LANGFUSE_BASE_URL    = "Langfuse Base URL"
	EnvDesc_LANGFUSE_PROJECT_ID  = "Langfuse Project ID"
	EnvDesc_LANGFUSE_PUBLIC_KEY  = "Langfuse Public Key"
	EnvDesc_LANGFUSE_SECRET_KEY  = "Langfuse Secret Key"

	// langfuse init variables
	EnvDesc_LANGFUSE_INIT_PROJECT_ID         = "Langfuse Init Project ID"
	EnvDesc_LANGFUSE_INIT_PROJECT_PUBLIC_KEY = "Langfuse Init Project Public Key"
	EnvDesc_LANGFUSE_INIT_PROJECT_SECRET_KEY = "Langfuse Init Project Secret Key"
	EnvDesc_LANGFUSE_INIT_USER_EMAIL         = "Langfuse Init User Email"
	EnvDesc_LANGFUSE_INIT_USER_NAME          = "Langfuse Init User Name"
	EnvDesc_LANGFUSE_INIT_USER_PASSWORD      = "Langfuse Init User Password"

	EnvDesc_LANGFUSE_OTEL_EXPORTER_OTLP_ENDPOINT = "Langfuse OTLP endpoint for OpenTelemetry exporter"

	EnvDesc_GRAFANA_LISTEN_IP     = "Grafana Listen IP"
	EnvDesc_GRAFANA_LISTEN_PORT   = "Grafana Listen Port"
	EnvDesc_OTEL_GRPC_LISTEN_IP   = "OTel gRPC Listen IP"
	EnvDesc_OTEL_GRPC_LISTEN_PORT = "OTel gRPC Listen Port"
	EnvDesc_OTEL_HTTP_LISTEN_IP   = "OTel HTTP Listen IP"
	EnvDesc_OTEL_HTTP_LISTEN_PORT = "OTel HTTP Listen Port"
	EnvDesc_OTEL_HOST             = "OpenTelemetry Host"

	EnvDesc_SUMMARIZER_PRESERVE_LAST       = "Summarizer Preserve Last"
	EnvDesc_SUMMARIZER_USE_QA              = "Summarizer Use QA"
	EnvDesc_SUMMARIZER_SUM_MSG_HUMAN_IN_QA = "Summarizer Human in QA"
	EnvDesc_SUMMARIZER_LAST_SEC_BYTES      = "Summarizer Last Section Bytes"
	EnvDesc_SUMMARIZER_MAX_BP_BYTES        = "Summarizer Max BP Bytes"
	EnvDesc_SUMMARIZER_MAX_QA_BYTES        = "Summarizer Max QA Bytes"
	EnvDesc_SUMMARIZER_MAX_QA_SECTIONS     = "Summarizer Max QA Sections"
	EnvDesc_SUMMARIZER_KEEP_QA_SECTIONS    = "Summarizer Keep QA Sections"

	EnvDesc_ASSISTANT_SUMMARIZER_PRESERVE_LAST    = "Assistant Summarizer Preserve Last"
	EnvDesc_ASSISTANT_SUMMARIZER_LAST_SEC_BYTES   = "Assistant Summarizer Last Section Bytes"
	EnvDesc_ASSISTANT_SUMMARIZER_MAX_BP_BYTES     = "Assistant Summarizer Max BP Bytes"
	EnvDesc_ASSISTANT_SUMMARIZER_MAX_QA_BYTES     = "Assistant Summarizer Max QA Bytes"
	EnvDesc_ASSISTANT_SUMMARIZER_MAX_QA_SECTIONS  = "Assistant Summarizer Max QA Sections"
	EnvDesc_ASSISTANT_SUMMARIZER_KEEP_QA_SECTIONS = "Assistant Summarizer Keep QA Sections"

	EnvDesc_EMBEDDING_PROVIDER        = "Embedding Provider"
	EnvDesc_EMBEDDING_URL             = "Embedding URL"
	EnvDesc_EMBEDDING_KEY             = "Embedding API Key"
	EnvDesc_EMBEDDING_MODEL           = "Embedding Model"
	EnvDesc_EMBEDDING_BATCH_SIZE      = "Embedding Batch Size"
	EnvDesc_EMBEDDING_STRIP_NEW_LINES = "Embedding Strip New Lines"
	EnvDesc_EMBEDDING_MAX_TEXT_BYTES  = "Embedding Max Text Bytes"

	EnvDesc_ASK_USER = "Human-in-the-loop"

	EnvDesc_ASSISTANT_USE_AGENTS = "Enable multi-agent mode for assistant"

	EnvDesc_EXECUTION_MONITOR_ENABLED          = "Enable Execution Monitoring (beta)"
	EnvDesc_EXECUTION_MONITOR_SAME_TOOL_LIMIT  = "Same Tool Call Threshold"
	EnvDesc_EXECUTION_MONITOR_TOTAL_TOOL_LIMIT = "Total Tool Call Threshold"
	EnvDesc_MAX_GENERAL_AGENT_TOOL_CALLS       = "Max Tool Calls for General Agents"
	EnvDesc_MAX_LIMITED_AGENT_TOOL_CALLS       = "Max Tool Calls for Limited Agents"
	EnvDesc_AGENT_PLANNING_STEP_ENABLED        = "Enable Task Planning (beta)"

	EnvDesc_SCRAPER_PUBLIC_URL                    = "Scraper Public URL"
	EnvDesc_SCRAPER_PRIVATE_URL                   = "Scraper Private URL"
	EnvDesc_LOCAL_SCRAPER_USERNAME                = "Local Scraper Username"
	EnvDesc_LOCAL_SCRAPER_PASSWORD                = "Local Scraper Password"
	EnvDesc_LOCAL_SCRAPER_MAX_CONCURRENT_SESSIONS = "Scraper Max Concurrent Sessions"

	EnvDesc_DUCKDUCKGO_ENABLED    = "DuckDuckGo Search"
	EnvDesc_DUCKDUCKGO_REGION     = "DuckDuckGo Region"
	EnvDesc_DUCKDUCKGO_SAFESEARCH = "DuckDuckGo Safe Search"
	EnvDesc_DUCKDUCKGO_TIME_RANGE = "DuckDuckGo Time Range"
	EnvDesc_SPLOITUS_ENABLED      = "Sploitus Search"
	EnvDesc_PERPLEXITY_API_KEY    = "Perplexity API Key"
	EnvDesc_TAVILY_API_KEY        = "Tavily API Key"
	EnvDesc_TRAVERSAAL_API_KEY    = "Traversaal API Key"
	EnvDesc_GOOGLE_API_KEY        = "Google Search API Key"
	EnvDesc_GOOGLE_CX_KEY         = "Google Search CX Key"
	EnvDesc_GOOGLE_LR_KEY         = "Google Search LR Key"

	EnvDesc_DOCKER_INSIDE                    = "Docker Inside Container"
	EnvDesc_DOCKER_NET_ADMIN                 = "Docker Network Admin"
	EnvDesc_DOCKER_SOCKET                    = "Docker Socket Path"
	EnvDesc_DOCKER_NETWORK                   = "Docker Network"
	EnvDesc_DOCKER_PUBLIC_IP                 = "Docker Public IP"
	EnvDesc_DOCKER_WORK_DIR                  = "Docker Work Directory"
	EnvDesc_DOCKER_DEFAULT_IMAGE             = "Docker Default Image"
	EnvDesc_DOCKER_DEFAULT_IMAGE_FOR_PENTEST = "Docker Pentest Image"
	EnvDesc_DOCKER_HOST                      = "Docker Host"
	EnvDesc_DOCKER_TLS_VERIFY                = "Docker TLS Verify"
	EnvDesc_DOCKER_CERT_PATH                 = "Docker Certificate Path"

	EnvDesc_LICENSE_KEY                       = "PentAGI License Key"
	EnvDesc_PENTAGI_LISTEN_IP                 = "PentAGI Server Host"
	EnvDesc_PENTAGI_LISTEN_PORT               = "PentAGI Server Port"
	EnvDesc_PUBLIC_URL                        = "PentAGI Public URL"
	EnvDesc_CORS_ORIGINS                      = "PentAGI CORS Origins"
	EnvDesc_COOKIE_SIGNING_SALT               = "PentAGI Cookie Signing Salt"
	EnvDesc_PROXY_URL                         = "HTTP/HTTPS Proxy URL"
	EnvDesc_HTTP_CLIENT_TIMEOUT               = "HTTP Client Timeout (seconds)"
	EnvDesc_TERMINAL_TOOL_TIMEOUT             = "Terminal Tool Timeout (seconds)"
	EnvDesc_EXTERNAL_SSL_CA_PATH              = "Custom CA Certificate Path"
	EnvDesc_EXTERNAL_SSL_INSECURE             = "Skip SSL Verification"
	EnvDesc_PENTAGI_SSL_DIR                   = "PentAGI SSL Directory"
	EnvDesc_PENTAGI_DATA_DIR                  = "PentAGI Data Directory"
	EnvDesc_PENTAGI_DOCKER_SOCKET             = "Mount Docker Socket Path"
	EnvDesc_PENTAGI_DOCKER_CERT_PATH          = "Mount Docker Certificate Path"
	EnvDesc_PENTAGI_LLM_SERVER_CONFIG_PATH    = "Custom LLM Host Config Path"
	EnvDesc_PENTAGI_OLLAMA_SERVER_CONFIG_PATH = "Ollama Host Config Path"

	EnvDesc_STATIC_DIR     = "Frontend Static Directory"
	EnvDesc_STATIC_URL     = "Frontend Static URL"
	EnvDesc_SERVER_PORT    = "Backend Server Port"
	EnvDesc_SERVER_HOST    = "Backend Server Host"
	EnvDesc_SERVER_SSL_CRT = "Backend Server SSL Certificate Path"
	EnvDesc_SERVER_SSL_KEY = "Backend Server SSL Key Path"
	EnvDesc_SERVER_USE_SSL = "Backend Server Use SSL"

	EnvDesc_PERPLEXITY_MODEL        = "Perplexity Model"
	EnvDesc_PERPLEXITY_CONTEXT_SIZE = "Perplexity Context Size"

	EnvDesc_SEARXNG_URL        = "Searxng Search URL"
	EnvDesc_SEARXNG_CATEGORIES = "Searxng Search Categories"
	EnvDesc_SEARXNG_LANGUAGE   = "Searxng Search Language"
	EnvDesc_SEARXNG_SAFESEARCH = "Searxng Safe Search"
	EnvDesc_SEARXNG_TIME_RANGE = "Searxng Time Range"
	EnvDesc_SEARXNG_TIMEOUT    = "Searxng Timeout"

	EnvDesc_OAUTH_GOOGLE_CLIENT_ID     = "OAuth Google Client ID"
	EnvDesc_OAUTH_GOOGLE_CLIENT_SECRET = "OAuth Google Client Secret"
	EnvDesc_OAUTH_GITHUB_CLIENT_ID     = "OAuth GitHub Client ID"
	EnvDesc_OAUTH_GITHUB_CLIENT_SECRET = "OAuth GitHub Client Secret"

	EnvDesc_LANGFUSE_EE_LICENSE_KEY   = "Langfuse Enterprise License Key"
	EnvDesc_PENTAGI_POSTGRES_PASSWORD = "PentAGI PostgreSQL Password"

	EnvDesc_GRAPHITI_URL        = "Graphiti Server URL"
	EnvDesc_GRAPHITI_TIMEOUT    = "Graphiti Request Timeout"
	EnvDesc_GRAPHITI_MODEL_NAME = "Graphiti Extraction Model"
	EnvDesc_NEO4J_USER          = "Neo4j Username"
	EnvDesc_NEO4J_DATABASE      = "Neo4j Database Name"
	EnvDesc_NEO4J_PASSWORD      = "Neo4j Database Password"
)

// dynamic, contextual sections used in processor operation forms
const (
	// section headers
	ProcessorSectionCurrentState = "Current state"
	ProcessorSectionPlanned      = "Planned actions"
	ProcessorSectionEffects      = "Effects"

	// component labels
	ProcessorComponentPentagi       = "PentAGI"
	ProcessorComponentLangfuse      = "Langfuse"
	ProcessorComponentObservability = "Observability"

	ProcessorComponentWorkerImage           = "worker image"
	ProcessorComponentComposeStacks         = "compose stacks"
	ProcessorComponentDefaultFiles          = "default files"
	ProcessorItemComposeFiles               = "compose files"
	ProcessorItemComposeStacksImagesVolumes = "compose stacks, images, volumes"

	// common states
	ProcessorStateInstalled = "installed"
	ProcessorStateMissing   = "not installed"
	ProcessorStateRunning   = "running"
	ProcessorStateStopped   = "stopped"
	ProcessorStateEmbedded  = "embedded"
	ProcessorStateExternal  = "external"
	ProcessorStateConnected = "connected"
	ProcessorStateDisabled  = "disabled"
	ProcessorStateUnknown   = "unknown"

	// planned action bullet prefixes
	PlannedWillStart    = "will start:"
	PlannedWillStop     = "will stop:"
	PlannedWillRestart  = "will restart:"
	PlannedWillUpdate   = "will update:"
	PlannedWillSkip     = "will skip:"
	PlannedWillRemove   = "will remove:"
	PlannedWillPurge    = "will purge:"
	PlannedWillDownload = "will download:"
	PlannedWillRestore  = "will restore:"

	// effect notes per operation (concise and practical)
	EffectsStart           = "PentAGI web UI becomes available. Background services are brought online in the required order."
	EffectsStop            = "Web UI becomes unavailable. In-progress flows pause safely. When you start PentAGI again, flows resume automatically. A small portion of the current agent step may be lost."
	EffectsRestart         = "Services stop and start again with a clean state. Brief downtime is expected. Flows resume automatically afterwards."
	EffectsUpdateAll       = "Images are pulled and services are recreated where needed. External or disabled components are skipped. Temporary downtime is expected."
	EffectsDownloadWorker  = "Running worker containers are not touched. New flows will use the downloaded image. To switch an existing flow to the new image, finish the flow and start a new task or create a new assistant."
	EffectsUpdateWorker    = "Pulls latest worker image. Running worker containers keep using the old image; new containers will use the updated one."
	EffectsUpdateInstaller = "The installer binary will be updated and the app will exit. Start the installer again to continue."
	EffectsFactoryReset    = "Removes containers, volumes and networks, restores default .env and embedded files. Produces a clean baseline. This action cannot be undone."
	EffectsRemove          = "Stops and removes containers but keeps volumes and images. Data is preserved. Web UI becomes unavailable until you start again."
	EffectsPurge           = "Complete cleanup: containers, images, volumes and configuration files are deleted. Irreversible."
	EffectsInstall         = "Required files are created and services are started. External components are detected and skipped."
)

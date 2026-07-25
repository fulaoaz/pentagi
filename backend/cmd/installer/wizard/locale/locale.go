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
	ChecksTitle               = "System Checks"
	ChecksWarningFailed       = "⚠ Some checks failed"
	CheckEnvironmentFile      = "Environment file"
	CheckWritePermissions     = "Write permissions"
	CheckDockerAPI            = "Docker API"
	CheckDockerVersion        = "Docker version"
	CheckDockerCompose        = "Docker Compose"
	CheckDockerComposeVersion = "Docker Compose version"
	CheckWorkerEnvironment    = "Worker environment"
	CheckSystemResources      = "System resources"
	CheckNetworkConnectivity  = "Network connectivity"
)

// EULA Screen constants
const (
	// Form interface implementation
	EULAFormDescription = "Legal terms and conditions for PentAGI usage"
	EULAFormName        = "EULA"
	EULAFormOverview    = `Review and accept the End User License Agreement to proceed with PentAGI installation.

The EULA contains:
• Software license terms and usage rights
• Limitation of liability and warranties
• Data collection and privacy policies
• Compliance requirements and restrictions
• Support and maintenance terms

You must scroll through the entire document and accept the terms to continue with the installation process.

Use arrow keys, page up/down, or home/end keys to navigate through the document.`

	// Error and status messages
	EULAErrorLoadingTitle     = "# Error Loading EULA\n\nFailed to load EULA: %v"
	EULAContentFallback       = "# EULA Content\n\n%s\n\n---\n\n*Note: Markdown rendering failed: %v*"
	EULAConfigurationRead     = "✓ EULA reviewed"
	EULAConfigurationAccepted = "✓ EULA accepted"
	EULAConfigurationPending  = "⚠ EULA not reviewed"
	EULALoading               = "Loading EULA..."
	EULAProgress              = "Progress: %d%%"
	EULAProgressComplete      = " • Complete"
)

// Main Menu Screen constants
const (
	MainMenuTitle       = "PentAGI Configuration"
	MainMenuDescription = "Configure all PentAGI components and settings"
	MainMenuName        = "Main Menu"
	MainMenuOverview    = `Welcome to PentAGI Configuration Center.

Configure essential components:
• LLM Providers - AI language models for autonomous testing
• Monitoring - Observability and analytics platforms
• Tools - Additional capabilities for enhanced testing
• System Settings - Environment and deployment options

Navigate through each section to complete your PentAGI setup.`

	MenuTitle        = "Configuration Menu"
	MenuSystemStatus = "System Status"
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
	LLMFormFieldAPIKey            = "API Key"
	LLMFormFieldDefaultAuth       = "Use Default AWS Auth"
	LLMFormFieldBearerToken       = "Bearer Token"
	LLMFormFieldAccessKey         = "Access Key ID"
	LLMFormFieldSecretKey         = "Secret Access Key"
	LLMFormFieldSessionToken      = "Session Token"
	LLMFormFieldRegion            = "Region"
	LLMFormFieldModel             = "Model"
	LLMFormFieldConfigPath        = "Config Path"
	LLMFormFieldLegacyReasoning   = "Legacy Reasoning"
	LLMFormFieldPreserveReasoning = "Preserve Reasoning"
	LLMFormFieldProviderName      = "Provider Name"
	LLMFormFieldPullTimeout       = "Model Pull Timeout"
	LLMFormFieldPullEnabled       = "Auto-pull Models"
	LLMFormFieldLoadModelsEnabled = "Load Models from Server"
	LLMFormBaseURLDesc            = "API endpoint URL for the provider"
	LLMFormAPIKeyDesc             = "Your API key for authentication"
	LLMFormDefaultAuthDesc        = "Use AWS SDK default credential chain (environment, EC2 role, ~/.aws/credentials) - highest priority"
	LLMFormBearerTokenDesc        = "Bearer token for authentication - takes priority over static credentials"
	LLMFormAccessKeyDesc          = "AWS Access Key ID for static credentials authentication"
	LLMFormSecretKeyDesc          = "AWS Secret Access Key for static credentials authentication"
	LLMFormSessionTokenDesc       = "AWS Session Token for temporary credentials (optional, used with static credentials)"
	LLMFormRegionDesc             = "AWS region for Bedrock service"
	LLMFormModelDesc              = "Default model to use for this provider"
	LLMFormConfigPathDesc         = "Path to configuration file (optional)"
	LLMFormLegacyReasoningDesc    = "Enable legacy reasoning mode (true/false)"
	LLMFormPreserveReasoningDesc  = "Preserve reasoning content in multi-turn conversations (required by some providers)"
	LLMFormProviderNameDesc       = "Provider name prefix for model names (useful for LiteLLM proxy)"
	LLMFormPullTimeoutDesc        = "Timeout in seconds for downloading models (default: 600)"
	LLMFormPullEnabledDesc        = "Automatically download required models on startup"
	LLMFormLoadModelsEnabledDesc  = "Load available models list from Ollama server"
	LLMFormOllamaAPIKeyDesc       = "Ollama Cloud API key (optional, leave empty for local Ollama server)"
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
	MonitoringLangfuseEmbedded = "Embedded Server"
	MonitoringLangfuseExternal = "External Server"
	MonitoringLangfuseDisabled = "Disabled"

	// Form fields
	MonitoringLangfuseDeploymentType     = "Deployment Type"
	MonitoringLangfuseDeploymentTypeDesc = "Select the deployment type for Langfuse"
	MonitoringLangfuseBaseURL            = "Server URL"
	MonitoringLangfuseBaseURLDesc        = "Address of the Langfuse server (e.g., https://cloud.langfuse.com)"
	MonitoringLangfuseProjectID          = "Project ID"
	MonitoringLangfuseProjectIDDesc      = "Project identifier in Langfuse"
	MonitoringLangfusePublicKey          = "Public Key"
	MonitoringLangfusePublicKeyDesc      = "Public API key for project access"
	MonitoringLangfuseSecretKey          = "Secret Key"
	MonitoringLangfuseSecretKeyDesc      = "Secret API key for project access"
	MonitoringLangfuseListenIP           = "Listen IP"
	MonitoringLangfuseListenIPDesc       = "Bind address used by Docker port mapping (e.g., 0.0.0.0 to expose on all interfaces)"
	MonitoringLangfuseListenPort         = "Listen Port"
	MonitoringLangfuseListenPortDesc     = "External TCP port exposed by Docker for Langfuse web UI"

	// Admin settings for embedded
	MonitoringLangfuseAdminEmail        = "Admin Email"
	MonitoringLangfuseAdminEmailDesc    = "Email for accessing the Langfuse admin panel"
	MonitoringLangfuseAdminPassword     = "Admin Password"
	MonitoringLangfuseAdminPasswordDesc = "Password for accessing the Langfuse admin panel"
	MonitoringLangfuseAdminName         = "Admin Username"
	MonitoringLangfuseAdminNameDesc     = "Administrator username in Langfuse"
	MonitoringLangfuseLicenseKey        = "Enterprise License Key"
	MonitoringLangfuseLicenseKeyDesc    = "Langfuse Enterprise license key (optional)"

	// Help text
	MonitoringLangfuseModeGuide    = "Choose deployment: Embedded (local control), External (cloud/existing), Disabled (no analytics)"
	MonitoringLangfuseEmbeddedHelp = `Embedded deploys complete Langfuse stack:
• PostgreSQL + ClickHouse databases
• MinIO S3 storage + Redis cache
• Full LLM conversation tracking
• Cost analysis and performance metrics
• Private data stays on your server

Resource requirements:
• ~2GB RAM, 5GB disk space minimum
• Additional storage for conversation logs
• Automatic setup and maintenance

Best for: Teams wanting data privacy, custom configurations, or no external dependencies. All analytics data stored locally with full administrative control.

Default admin access:
• Web UI: http://localhost:4000
• Login: admin@pentagi.com
• Password: password (change required)`
	MonitoringLangfuseExternalHelp = `External connects to cloud.langfuse.com or your existing Langfuse server:

• No local infrastructure needed
• Managed updates and maintenance
• Shared analytics across teams
• Enterprise features available
• Data stored on external provider

Setup requirements:
• Langfuse account and API keys
• Internet connectivity required
• Project ID and authentication keys

Best for: Teams using cloud services, wanting managed infrastructure, or integrating with existing Langfuse deployments across organizations.`
	MonitoringLangfuseDisabledHelp = `Langfuse is disabled. Without LLM observability you will not have:

• Conversation history tracking
• Token usage and cost analysis
• Model performance metrics
• Debug traces for AI interactions
• User behavior analytics
• Prompt engineering insights

Consider enabling for production use
to monitor AI agent performance and
optimize costs effectively.`
)

// Graphiti Integration constants
const (
	MonitoringGraphitiFormTitle       = "Graphiti Configuration (beta)"
	MonitoringGraphitiFormDescription = "Configuration of Graphiti knowledge graph integration"
	MonitoringGraphitiFormName        = "Graphiti (beta)"
	MonitoringGraphitiFormOverview    = `⚠️  BETA FEATURE: This functionality is currently under active development. Please monitor updates for improvements and stability fixes.

Graphiti provides temporal knowledge graph capabilities:
• Entity and relationship extraction
• Semantic memory for AI agents
• Temporal context tracking
• Knowledge reuse across flows

⚠️  REQUIREMENT: Graphiti requires configured OpenAI provider (LLM Providers → OpenAI) for entity extraction.

Choose between embedded instance or external connection.`

	// Deployment types
	MonitoringGraphitiEmbedded = "Embedded Stack"
	MonitoringGraphitiExternal = "External Service"
	MonitoringGraphitiDisabled = "Disabled"

	// Form fields
	MonitoringGraphitiDeploymentType     = "Deployment Type"
	MonitoringGraphitiDeploymentTypeDesc = "Select the deployment type for Graphiti"
	MonitoringGraphitiURL                = "Graphiti Server URL"
	MonitoringGraphitiURLDesc            = "Address of the Graphiti API server"
	MonitoringGraphitiTimeout            = "Request Timeout"
	MonitoringGraphitiTimeoutDesc        = "Timeout in seconds for Graphiti operations"
	MonitoringGraphitiModelName          = "Extraction Model"
	MonitoringGraphitiModelNameDesc      = "LLM model for entity extraction (uses OpenAI provider from LLM Providers configuration)"
	MonitoringGraphitiNeo4jUser          = "Neo4j Username"
	MonitoringGraphitiNeo4jUserDesc      = "Username for Neo4j database access"
	MonitoringGraphitiNeo4jPassword      = "Neo4j Password"
	MonitoringGraphitiNeo4jPasswordDesc  = "Password for Neo4j database access"
	MonitoringGraphitiNeo4jDatabase      = "Neo4j Database"
	MonitoringGraphitiNeo4jDatabaseDesc  = "Neo4j database name"

	// Help text
	MonitoringGraphitiModeGuide    = "Choose deployment: Embedded (local Neo4j), External (existing Graphiti), Disabled (no knowledge graph)"
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
	SummarizerTitle       = "Summarizer Configuration"
	SummarizerDescription = "Enable conversation summarization to reduce LLM costs and improve context management"
	SummarizerName        = "Summarizer"
	SummarizerOverview    = `Optimize context usage, reduce LLM costs, and match your model capabilities.

When to adjust summarization:
• High token costs: Reduce context size (4K-12K vs 22K+ tokens)
• "Context too long" errors: Configure for your model's limits
• Poor conversation flow: Increase context retention for quality
• Different model types: Short-context vs long-context model tuning

General Summarization: Maximum cost control and precision tuning for research/analysis tasks
Assistant Summarization: Optimal conversation quality with intelligent context management for interactive sessions

Quick wins:
• Cost reduction: Use General, reduce Recent Sections to 1-2
• Context errors: Match limits to your model (8K/32K/128K)
• Quality priority: Use Assistant with increased limits`

	SummarizerTypeGeneralName = "General Summarization"
	SummarizerTypeGeneralDesc = "Global summarization settings for conversation context management"

	SummarizerTypeGeneralInfo = `Choose this for maximum cost control and short-context model compatibility.

Perfect when you need:
• Aggressive cost reduction: Fine-tune every parameter for minimal token usage
• Short-context models (8K-32K): Precise limits to avoid overflow errors
• Research/analysis tasks: Controlled compression without losing key data
• Custom QA handling: Full control over question-answer pair processing

Typical results:
• 40-70% cost reduction vs default settings
• 4K-12K token contexts (vs 22K+ in Assistant mode)
• Better performance on GPT-3.5, Claude Instant, smaller models
• Precise control over conversation memory vs fresh context balance

Best practices:
• Start with 1-2 Recent Sections for maximum savings
• Enable Size Management for automatic overflow protection
• Disable QA compression only for critical reasoning tasks`

	SummarizerTypeAssistantName = "Assistant Summarization"
	SummarizerTypeAssistantDesc = "Specialized summarization settings for AI assistant contexts"

	SummarizerTypeAssistantInfo = `Choose this for optimal conversation quality and dialogue continuity.

Perfect when you need:
• Extended reasoning chains: Maintain context for complex multi-step thinking
• High-quality conversations: Preserve dialogue flow and assistant personality
• Long-context models (64K+): Leverage full model capabilities efficiently
• Interactive sessions: Better memory of user preferences and conversation history

Typical results:
• 8K-40K token contexts with intelligent compression
• Superior conversation continuity vs manual settings
• Automatic context optimization for reasoning tasks
• Balanced cost vs quality (3x more context than General mode)

Best practices:
• Use default settings for most scenarios - they're pre-optimized
• Increase Recent Sections only for very complex tasks
• Monitor context usage - costs scale with token count
• Perfect for GPT-4, Claude, and other large context models`
)

// Summarizer Form Screen
const (
	SummarizerFormGeneralTitle   = "General Summarizer Configuration"
	SummarizerFormAssistantTitle = "Assistant Summarizer Configuration"
	SummarizerFormDescription    = "Configure %s Settings"

	// Field Labels and Descriptions
	SummarizerFormPreserveLast     = "Size Management"
	SummarizerFormPreserveLastDesc = "Controls last section compression. Enabled: sections fit LastSecBytes (smaller context). Disabled: sections grow freely (larger context)"

	SummarizerFormUseQA     = "QA Summarization"
	SummarizerFormUseQADesc = "Enables question-answer pair compression when total QA content exceeds MaxQABytes or MaxQASections limits"

	SummarizerFormSumHumanInQA     = "Compress User Messages"
	SummarizerFormSumHumanInQADesc = "Include user messages in QA compression. Disabled: preserves original user text (recommended for most cases)"

	SummarizerFormLastSecBytes     = "Section Size Limit"
	SummarizerFormLastSecBytesDesc = "Maximum bytes per recent section when Size Management enabled. Larger: more detail per section, higher token usage"

	SummarizerFormMaxBPBytes     = "Response Size Limit"
	SummarizerFormMaxBPBytesDesc = "Maximum bytes for individual AI responses before compression. Prevents single large responses from dominating context"

	SummarizerFormMaxQASections     = "QA Section Limit"
	SummarizerFormMaxQASectionsDesc = "Maximum question-answer sections before QA compression triggers. Works with MaxQABytes to control total QA memory"

	SummarizerFormMaxQABytes     = "Total QA Memory"
	SummarizerFormMaxQABytesDesc = "Maximum bytes for all QA sections combined. When exceeded (with MaxQASections), triggers QA compression to fit limit"

	SummarizerFormKeepQASections     = "Recent Sections"
	SummarizerFormKeepQASectionsDesc = "Number of most recent conversation sections preserved without compression. PRIMARY parameter affecting context size"

	// Enhanced Help Text - General (common principles)
	SummarizerFormGeneralHelp = `Context estimation: 4K-22K tokens (typical), up to 94K (maximum settings).

Key relationships:
• Recent Sections: Most critical - each +1 adds ~1.5-9K tokens
• Size Management OFF: 2-3x larger context (less compression)
• Section/Response Limits: Control individual component sizes
• QA Memory: Manages total conversation history when limits exceeded

Parameter interactions:
• QA compression activates when BOTH MaxQABytes AND MaxQASections exceeded
• Size Management disabled → sections can grow 2x larger than limits
• Response Limit prevents single large outputs from dominating context
• User message compression (SummHumanInQA) saves 5% but loses original phrasing

Reduce for smaller models:
• Recent Sections: 1-2 (vs 3+ default)
• Section Limit: 25-35KB (vs 50KB+)
• Disable Size Management for simple conversations

Common mistakes:
• Setting Recent Sections too high (main cause of context overflow)
• Enabling Size Management with very low Section Limits (over-compression)
• Mismatched QA limits (high bytes + low sections = ineffective)

Current algorithm compresses older content while preserving recent context quality.`

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
	ToolsTitle       = "Tools Configuration"
	ToolsDescription = "Enhance agent capabilities with additional tools and options"
	ToolsName        = "Tools"
	ToolsOverview    = `Configure additional tools and capabilities for AI agents.
Each tool can be enabled and configured according to your requirements.

Available settings:
• Human-in-the-loop - Enable user interaction during testing
• AI Agents Settings - Configure global behavior for AI agents
• Search Engines - Configure external search providers
• Scraper - Web content extraction and analysis
• Graphiti (beta) - Temporal knowledge graph for semantic memory
• Docker - Container environment configuration`
)

// Server Settings screen strings
const (
	ServerSettingsFormTitle       = "Server Settings"
	ServerSettingsFormDescription = "Configure PentAGI server network access and public routing"
	ServerSettingsFormName        = "Server Settings"
	ServerSettingsFormOverview    = `• Network binding - control which interface and port PentAGI listens on
• Public URL - external address and optional base path used in redirects
• CORS - allowed origins for browser access
• Proxy - HTTP/HTTPS proxy for outbound traffic to LLM/search providers
• SSL directory - custom certificates directory containing server.crt and server.key (PEM)
• Data directory - persistent storage for agent artifacts and flow workspaces`

	// Field labels and descriptions
	ServerSettingsLicenseKey     = "License Key"
	ServerSettingsLicenseKeyDesc = "PentAGI License Key in format of XXXX-XXXX-XXXX-XXXX"

	ServerSettingsHost     = "Server Host (Listen IP)"
	ServerSettingsHostDesc = "Bind address used by Docker port mapping (e.g., 0.0.0.0 to expose on all interfaces)"

	ServerSettingsPort     = "Server Port (Listen Port)"
	ServerSettingsPortDesc = "External TCP port exposed by Docker for PentAGI web UI"

	ServerSettingsPublicURL     = "Public URL"
	ServerSettingsPublicURLDesc = "Base public URL for redirects and links (supports base path, e.g., https://example.com/pentagi/)"

	ServerSettingsCORSOrigins     = "CORS Origins"
	ServerSettingsCORSOriginsDesc = "Comma-separated list of allowed origins (e.g., https://localhost:8443,https://localhost)"

	ServerSettingsProxyURL     = "HTTP/HTTPS Proxy"
	ServerSettingsProxyURLDesc = "Proxy for outbound requests to LLMs and external tools (not used for Docker API access)"

	ServerSettingsProxyUsername     = "Proxy Username"
	ServerSettingsProxyUsernameDesc = "Username for proxy authentication (optional)"
	ServerSettingsProxyPassword     = "Proxy Password"
	ServerSettingsProxyPasswordDesc = "Password for proxy authentication (optional)"

	ServerSettingsHTTPClientTimeout       = "HTTP Client Timeout"
	ServerSettingsHTTPClientTimeoutDesc   = "Timeout in seconds for external API calls (LLM providers, search engines, etc.)"
	ServerSettingsTerminalToolTimeout     = "Terminal Tool Timeout"
	ServerSettingsTerminalToolTimeoutDesc = "Default timeout in seconds for terminal commands (0 or negative = use 3-hour maximum)"

	ServerSettingsExternalSSLCAPath     = "Custom CA Certificate Path"
	ServerSettingsExternalSSLCAPathDesc = "Path inside container to custom root CA cert (e.g., /opt/pentagi/ssl/ca-bundle.pem)"

	ServerSettingsExternalSSLInsecure     = "Skip SSL Verification"
	ServerSettingsExternalSSLInsecureDesc = "Disable SSL/TLS certificate validation (use only for testing with self-signed certs)"

	ServerSettingsSSLDir     = "SSL Directory"
	ServerSettingsSSLDirDesc = "Directory containing server.crt and server.key in PEM format (server.crt may include fullchain)"

	ServerSettingsDataDir     = "Data Directory"
	ServerSettingsDataDirDesc = "Directory for all agent-generated files; contains flow-N subdirectories used as /work in worker containers"

	ServerSettingsCookieSigningSalt     = "Cookie Signing Salt"
	ServerSettingsCookieSigningSaltDesc = "Secret used to sign cookies (keep private)"

	// Hints for fields overview
	ServerSettingsLicenseKeyHint          = "License Key"
	ServerSettingsHostHint                = "Listen IP"
	ServerSettingsPortHint                = "Listen Port"
	ServerSettingsPublicURLHint           = "Public URL"
	ServerSettingsCORSOriginsHint         = "CORS Origins"
	ServerSettingsProxyURLHint            = "Proxy URL"
	ServerSettingsProxyUsernameHint       = "Proxy Username"
	ServerSettingsProxyPasswordHint       = "Proxy Password"
	ServerSettingsHTTPClientTimeoutHint   = "HTTP Timeout"
	ServerSettingsTerminalToolTimeoutHint = "Terminal Timeout"
	ServerSettingsExternalSSLCAPathHint   = "Custom CA Path"
	ServerSettingsExternalSSLInsecureHint = "Skip SSL Verification"
	ServerSettingsSSLDirHint              = "SSL Directory"
	ServerSettingsDataDirHint             = "Data Directory"

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
	ToolsAIAgentsSettingsFormTitle       = "AI Agents Settings"
	ToolsAIAgentsSettingsFormDescription = "Configure global behavior for AI agents"
	ToolsAIAgentsSettingsFormName        = "AI Agents Settings"
	ToolsAIAgentsSettingsFormOverview    = `This section configures global behavior of AI agents across PentAGI.

Basic Settings:
• Enable User Interaction: allow agents to request user input when needed
• Use Multi-Agent Mode: enable assistant to orchestrate multiple specialized agents

Execution Monitoring (⚠️  BETA):
• Enable Execution Monitoring: automatic mentor supervision for pattern analysis
• Same Tool Call Threshold: consecutive identical tool calls before mentor review
• Total Tool Call Threshold: total tool calls before mentor review

Tool Call Limits:
• Max Tool Calls (General Agents): prevent runaway executions for Assistant, Primary Agent, Pentester, Coder, Installer
• Max Tool Calls (Limited Agents): prevent runaway executions for Searcher, Enricher, Memorist, etc.

Task Planning (⚠️  BETA):
• Enable Task Planning: generate structured execution plans for specialist agents

⚠️  BETA features are under active development. Enable for testing only.`

	// field labels and descriptions
	ToolsAIAgentsSettingHumanInTheLoop          = "Enable User Interaction"
	ToolsAIAgentsSettingHumanInTheLoopDesc      = "Allow agents to ask for user input when needed"
	ToolsAIAgentsSettingUseAgents               = "Use Multi-Agent Mode"
	ToolsAIAgentsSettingUseAgentsDesc           = "Enable assistant to orchestrate multiple specialized agents"
	ToolsAIAgentsSettingExecutionMonitor        = "Enable Execution Monitoring (beta)"
	ToolsAIAgentsSettingExecutionMonitorDesc    = "Automatically invoke mentor for execution pattern analysis"
	ToolsAIAgentsSettingSameToolLimit           = "Same Tool Call Threshold"
	ToolsAIAgentsSettingSameToolLimitDesc       = "Consecutive identical tool calls before mentor review"
	ToolsAIAgentsSettingTotalToolLimit          = "Total Tool Call Threshold"
	ToolsAIAgentsSettingTotalToolLimitDesc      = "Total tool calls before mentor review"
	ToolsAIAgentsSettingMaxGeneralToolCalls     = "Max Tool Calls (General Agents)"
	ToolsAIAgentsSettingMaxGeneralToolCallsDesc = "Maximum tool calls for Assistant, Primary Agent, Pentester, Coder, Installer"
	ToolsAIAgentsSettingMaxLimitedToolCalls     = "Max Tool Calls (Limited Agents)"
	ToolsAIAgentsSettingMaxLimitedToolCallsDesc = "Maximum tool calls for Searcher, Enricher, Memorist, etc."
	ToolsAIAgentsSettingTaskPlanning            = "Enable Task Planning (beta)"
	ToolsAIAgentsSettingTaskPlanningDesc        = "Generate structured execution plans for specialist agents"

	// help content
	ToolsAIAgentsSettingsHelp = `AI Agents Settings define how agents collaborate, interact with users, and handle execution control.

Basic Settings:
• Enable User Interaction: allow agents to ask for user input when needed
• Use Multi-Agent Mode: enable assistant to orchestrate specialized agents for complex tasks

Execution Monitoring (⚠️  BETA):
Automatically invokes adviser (mentor) to analyze execution patterns, detect loops, suggest alternative strategies, and prevent agents from fixating on single approach. Thresholds: consecutive identical calls (default: 5) and total calls (default: 10).

Task Planning (⚠️  BETA):
Generates 3-7 step execution plans before specialist agents begin work. Prevents scope creep and improves success rates. Works best when adviser uses enhanced configuration (stronger model or maximum reasoning mode).

Tool Call Limits (always active):
Hard limits prevent infinite loops: General agents default 100, Limited agents default 20. Works independently from beta features.

OPEN SOURCE MODELS < 32B (Qwen3.5-27B, DeepSeek-V3, Llama-3.1-70B):
✓ ENABLE both beta features - ESSENTIAL for quality results
✓ Testing shows 2x improvement in result quality vs. baseline
✓ Configure adviser with enhanced settings for best performance
✓ Ideal for air-gapped deployments with local LLM inference

Performance: 2-3x increase in tokens/time, 2x improvement in quality for models < 32B.

⚠️  BETA WARNING: Features under active development. Recommended for open source models < 32B despite beta status. For cloud APIs with larger models, keep disabled.

Note: Changes require service restart.`
)

// Search Engines screen strings
const (
	ToolsSearchEnginesFormTitle       = "Search Engines Configuration"
	ToolsSearchEnginesFormDescription = "Configure search engines for AI agents to gather intelligence during testing"
	ToolsSearchEnginesFormName        = "Search Engines"
	ToolsSearchEnginesFormOverview    = `Available search engines:
• DuckDuckGo - Free search engine (no API key required)
• Sploitus - Security exploits and vulnerabilities database (no API key required)
• Perplexity - AI-powered search with reasoning
• Tavily - Search API for AI applications
• Traversaal - Web scraping and search
• Google Search - Requires API key and Custom Search Engine ID
• Searxng - Internet metasearch engine

Get API keys from:
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

Host Mode (value: 'host'):
• Direct access to host network interfaces
• No port forwarding - ports bind directly to host
• Required for raw packet manipulation
• Advanced network testing capabilities
• Lower isolation - use with caution

Examples:
• 'pentagi-network' - creates isolated bridge network
• 'host' - enables direct host network access

Security Note: Host network mode reduces container isolation. Only use when necessary for advanced penetration testing tasks requiring direct network stack access.`

	ToolsDockerPublicIPHelp = `Public IP Address enables out-of-band (OOB) attack techniques by providing workers with a reachable address for reverse connections.

Workers automatically receive two random ports (28000-32000 range) mapped to this IP for receiving callbacks from exploited targets.

By default agents will try to get public address from the services api.ipify.org, ipinfo.io/ip or ifconfig.me.`

	ToolsDockerWorkDirHelp = `Work Directory specifies host filesystem location for worker storage. When set, replaces default Docker volumes with host directory mounts.

Benefits:
• Persistent storage across restarts
• Direct file system access
• Easier artifact management
• Custom backup strategies

By default uses Docker dedicated volume per worker container.

Example: /path/to/workdir/`

	ToolsDockerDefaultImageHelp = `Default Image provides fallback for workers when task requirements don't specify a particular container image.

Should contain basic utilities and tools for general-purpose tasks. Default: debian:latest`

	ToolsDockerDefaultImageForPentestHelp = `Pentesting Image serves as default for security testing tasks. Should include comprehensive security tools and utilities.

Recommended images include Kali Linux, Parrot Security, or custom security-focused containers. Default: vxcontrol/kali-linux`

	ToolsDockerHostHelp = `Docker Host uses for start primary worker containers and overrides default Docker daemon connection. Supports Unix sockets and TCP connections.

Examples:
• unix:///var/run/docker.sock (local)
• tcp://docker-host:2376 (remote)

Enable TLS for remote connections.`

	ToolsDockerTLSVerifyHelp = `TLS Verification secures Docker daemon connections over TCP. Strongly recommended for remote Docker hosts.

Requires valid certificates in the specified certificate directory.`

	ToolsDockerCertPathHelp = `TLS Certificates directory must contain:
• ca.pem - Certificate Authority
• cert.pem - Client certificate
• key.pem - Private key

Required for secure remote Docker connections when using TLS to manage worker containers.

Example: /path/to/certs`
)

// Embedder form strings
const (
	EmbedderFormTitle       = "Embedder Configuration"
	EmbedderFormDescription = "Configure text vectorization for semantic search and knowledge storage"
	EmbedderFormName        = "Embedder"
	EmbedderFormOverview    = `Text embeddings convert documents into vectors for semantic search and knowledge storage.
Different providers offer various models with different capabilities and pricing.

Choose carefully as changing providers requires reindexing all stored data.`

	EmbedderFormProvider     = "Embedding Provider"
	EmbedderFormProviderDesc = "Select the provider for text vectorization. Embeddings are used for semantic search and knowledge storage."

	EmbedderFormURL     = "API Endpoint URL"
	EmbedderFormURLDesc = "Custom API endpoint (leave empty to use default)"

	EmbedderFormAPIKey     = "API Key"
	EmbedderFormAPIKeyDesc = "Authentication key for the provider (not required for Ollama)"

	EmbedderFormModel     = "Model Name"
	EmbedderFormModelDesc = "Specific embedding model to use (leave empty for provider default)"

	EmbedderFormBatchSize     = "Batch Size"
	EmbedderFormBatchSizeDesc = "Number of documents to process in a single batch (1-1000)"

	EmbedderFormStripNewLines     = "Strip New Lines"
	EmbedderFormStripNewLinesDesc = "Remove line breaks from text before embedding (true/false)"

	EmbedderFormMaxTextBytes     = "Max Text Bytes"
	EmbedderFormMaxTextBytesDesc = "Maximum number of bytes per text chunk sent to the embedding API (e.g. 8192)"

	EmbedderFormHelpTitle   = "Embedding Configuration"
	EmbedderFormHelpContent = `Configure text vectorization for semantic search and knowledge storage.

If no specific embedding settings are configured, the system will use OpenAI embeddings with the API key from LLM Providers.

Change providers carefully - different embedders produce incompatible vectors requiring database reindexing.`

	EmbedderFormHelpOpenAI      = "OpenAI: Most reliable option with excellent quality. Requires API key from LLM Providers if not set here."
	EmbedderFormHelpOllama      = "Ollama: Local embeddings, no API key needed. Requires Ollama server running."
	EmbedderFormHelpHuggingFace = "HuggingFace: Open source models with API key required."
	EmbedderFormHelpGoogleAI    = "Google AI: Quality embeddings, requires API key."

	// Provider names and descriptions
	EmbedderProviderDefault         = "Default (OpenAI)"
	EmbedderProviderDefaultDesc     = "Use OpenAI embeddings with API key from LLM Providers configuration"
	EmbedderProviderOpenAI          = "OpenAI"
	EmbedderProviderOpenAIDesc      = "OpenAI text embeddings API (text-embedding-3-small, ada-002)"
	EmbedderProviderOllama          = "Ollama"
	EmbedderProviderOllamaDesc      = "Local Ollama server for open-source embedding models"
	EmbedderProviderMistral         = "Mistral"
	EmbedderProviderMistralDesc     = "Mistral AI embedding models"
	EmbedderProviderJina            = "Jina"
	EmbedderProviderJinaDesc        = "Jina AI embedding API"
	EmbedderProviderHuggingFace     = "HuggingFace"
	EmbedderProviderHuggingFaceDesc = "HuggingFace inference API for embedding models"
	EmbedderProviderGoogleAI        = "Google AI"
	EmbedderProviderGoogleAIDesc    = "Google AI embedding models (embedding-001)"
	EmbedderProviderVoyageAI        = "VoyageAI"
	EmbedderProviderVoyageAIDesc    = "VoyageAI embedding API"
	EmbedderProviderDisabled        = "Disabled"
	EmbedderProviderDisabledDesc    = "Disable embeddings functionality completely"

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
	MockScreenTitle       = "Development Screen"
	MockScreenDescription = "This screen is under development"
)

// Apply Changes screen constants
const (
	ApplyChangesFormTitle       = "Apply Configuration Changes"
	ApplyChangesFormName        = "Apply Changes"
	ApplyChangesFormDescription = "Review and apply your configuration changes"

	// Apply Changes overview and help
	ApplyChangesFormOverview = `This screen allows you to review all pending configuration changes and apply them to your PentAGI installation.

When you apply changes, the system will:
• Save all modified environment variables to the .env file
• Restart affected services with the new configuration
• Install additional components if needed`

	// Apply Changes status messages
	ApplyChangesNotStarted     = "Configuration changes are ready to be applied"
	ApplyChangesInProgress     = "Applying configuration changes...\n"
	ApplyChangesCompleted      = "Configuration changes have been successfully applied\n"
	ApplyChangesFailed         = "Failed to perform configuration changes"
	ApplyChangesResetCompleted = "Configuration changes have been successfully reset\n"

	ApplyChangesTerminalIsNotInitialized = "Terminal is not initialized"

	// Apply Changes instructions
	ApplyChangesInstructions = `Press Enter to begin applying the configuration changes.`

	ApplyChangesNoChanges = "No configuration changes are pending"

	// Apply Changes installation status
	ApplyChangesInstallNotFound = `PentAGI is not currently installed on this system.

The following actions will be performed:
• Docker environment setup and validation
• Creation of docker-compose.yml file
• Installation and startup of PentAGI core services`

	ApplyChangesInstallFoundLangfuse      = `• Installation of Langfuse observability stack (docker-compose-langfuse.yml)`
	ApplyChangesInstallFoundObservability = `• Installation of comprehensive observability stack with Grafana, VictoriaMetrics, and Jaeger (docker-compose-observability.yml)`

	ApplyChangesUpdateFound = `PentAGI is currently installed on this system.

The following actions will be performed:
• Update environment variables in .env file
• Recreate and restart affected Docker containers
• Apply new configuration to running services`

	// Apply Changes warnings and notes
	ApplyChangesWarningCritical = "⚠️  Critical changes detected - services will be restarted"
	ApplyChangesWarningSecrets  = "🔒 Secret values detected - they will be securely stored"
	ApplyChangesNoteBackup      = "💾 Current configuration will be backed up before changes"
	ApplyChangesNoteTime        = "⏱️  This process may take less than a minute depending on selected components"

	// Apply Changes progress messages
	ApplyChangesStageValidation = "Validating environment and dependencies..."
	ApplyChangesStageBackup     = "Creating configuration backup..."
	ApplyChangesStageEnvFile    = "Updating environment file..."
	ApplyChangesStageCompose    = "Generating Docker Compose files..."
	ApplyChangesStageDocker     = "Managing Docker containers..."
	ApplyChangesStageServices   = "Starting services..."
	ApplyChangesStageComplete   = "Configuration changes applied successfully"

	// Apply Changes change list headers
	ApplyChangesChangesTitle  = "Pending Configuration Changes"
	ApplyChangesChangesCount  = "Total changes: %d"
	ApplyChangesChangesMasked = "(hidden for security)"
	ApplyChangesChangesEmpty  = "No changes to apply"

	// Apply Changes help content
	ApplyChangesHelpTitle   = "Applying Configuration Changes"
	ApplyChangesHelpContent = `Be sure to check the current configuration before applying changes.`
)

// apply changes integrity prompt
const (
	ApplyChangesIntegrityPromptTitle   = "File integrity check"
	ApplyChangesIntegrityPromptMessage = "Out-of-date files were detected.\nDo you want to update them to the latest version?"
	ApplyChangesIntegrityOutdatedList  = "Out-of-date files:\n%s\nConfirm update? (y/n)"
	ApplyChangesIntegrityChecking      = "Collecting file integrity information..."
	ApplyChangesIntegrityNoOutdated    = "No out-of-date files found. Proceeding with apply."
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

# PentAGI

<div align="center" style="font-size: 1.5em; margin: 20px 0;">
    <strong>P</strong>enetration testing <strong>A</strong>rtificial <strong>G</strong>eneral <strong>I</strong>ntelligence
</div>

<div align="center">

[简体中文](README.md) | [English](README.en.md) | [汉化维护与上游同步](UPSTREAM_SYNC.zh-CN.md)

</div>

## 汉化版说明

这是 PentAGI 的简体中文维护分支，默认分支为 `zh-CN`。汉化以用户实际可见、需要理解和操作的内容为重点，包括网页界面、交互式安装器、REST/GraphQL 错误提示，以及项目自有的 Grafana 面板。仓库内 `backend/cmd/installer` 的终端安装器也已完成汉化，并由 GitHub Actions 自动构建多平台版本。

项目发布说明、界面预览与部署体验见：[浮潦の小窝 - PentAGI 简体中文维护版发布](https://fulao.cc/2026/07/27/pentagi-zh-cn/)。

以下内容有意保留英文，避免改变程序契约、影响模型效果或增加无意义的上游冲突：

- PentAGI 品牌、模型 ID、命令、环境变量名、协议字段和枚举实际值。
- AI 系统提示词、工具调用协议、工具原始输出和服务端技术日志。
- 自动生成的 Swagger 文档，以及直接引入的第三方 Grafana 面板。

本分支通过 GitHub Actions 每小时检查 `vxcontrol/pentagi:main`。发现更新后会创建同步 PR，先经过汉化门禁和 CI 复核；合并到 `zh-CN` 后，会自动构建并发布以下 Linux amd64 镜像：

- `ghcr.io/fulaoaz/pentagi:zh-cn`：最新汉化版。
- `ghcr.io/fulaoaz/pentagi:zh-cn-<commit>`：按提交锁定的可回滚版本。

与上游安装方式相比，本分支的 `docker-compose.yml`、`.env.example` 和汉化安装器默认使用上述 GHCR 汉化镜像。汉化安装器可从 [`zh-cn-installer-latest`](https://github.com/fulaoaz/pentagi/releases/tag/zh-cn-installer-latest) 下载；`pentagi.com` 提供的安装器属于上游官方发行渠道，不是本分支构建的汉化版本。

当前版本已在 Windows 11、WSL2 Ubuntu 24.04 和 Docker Desktop 环境完成实际验证：镜像拉取、数据库迁移、HTTPS 访问、默认管理员登录、首次改密流程和中文仪表盘均正常。首次访问本机地址时，浏览器会提示自签名证书，这是默认本地 TLS 配置的预期行为。

### 汉化界面预览

![PentAGI 简体中文仪表盘](docs/images/zh-cn-dashboard.jpg)

> 截图来自 `zh-CN` 分支提交 `3101d2b` 构建的 GHCR 镜像，并在本机 WSL2 部署后通过默认管理员登录实测。图中展示了中文导航、数据分析页签、时间范围、图表说明和空状态。

> [!IMPORTANT]
> PentAGI 是安全测试项目，仓库中的 Sploitus 测试夹具会保存公开漏洞样本文本。例如 `backend/pkg/tools/testdata/sploitus_result_nginx.json` 可能被 Windows Defender 按其中的 PHP WebShell 特征识别为后门，但它本身是不会被项目执行的 JSON 测试数据。遇到告警时应核对命中路径、文件类型和 Defender 的“是否执行”状态，不要直接为整个仓库添加杀毒排除项；仓库外的可执行文件告警也应单独处置。
<br>
<div align="center">

> **加入社区！** 与安全研究人员、AI 爱好者和志同道合的白帽黑客交流，获取支持、分享经验，并及时了解 PentAGI 的最新进展。

[![Discord](https://img.shields.io/badge/Discord-7289DA?logo=discord&logoColor=white)](https://discord.gg/2xrMh7qX6m)⠀[![Telegram](https://img.shields.io/badge/Telegram-2CA5E0?logo=telegram&logoColor=white)](https://t.me/+Ka9i6CNwe71hMWQy)

<a href="https://trendshift.io/repositories/15161" target="_blank"><img src="https://trendshift.io/api/badge/repositories/15161" alt="vxcontrol%2Fpentagi | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>

</div>

## 目录

- [汉化版说明](#汉化版说明)
- [概述](#概述)
- [功能特性](#功能特性)
- [架构](#架构)
  - [智能体监管](#advanced-agent-supervision)
- [快速开始](#快速开始)
  - [使用汉化镜像（推荐）](#使用汉化镜像推荐)
  - [使用汉化安装程序](#使用汉化安装程序)
  - [手动安装](#手动安装)
- [登录后如何使用 PentAGI](#登录后如何使用-pentagi)
- [API 访问](#api-访问)
  - [LLM 提供商配置](#自定义-llm-提供商配置)
    - [Ollama](#ollama-提供商配置)
    - [OpenAI](#openai-提供商配置)
    - [Anthropic](#anthropic-提供商配置)
    - [Google AI（Gemini）](#google-aigemini提供商配置)
    - [AWS Bedrock](#aws-bedrock-提供商配置)
    - [DeepSeek](#deepseek-提供商配置)
    - [GLM](#glm-提供商配置)
    - [Kimi](#kimi-提供商配置)
    - [Qwen](#qwen-提供商配置)
- [高级设置](#高级设置)
  - [集成 Langfuse](#集成-langfuse)
  - [监控与可观测性](#监控与可观测性)
  - [集成知识图谱（Graphiti）](#集成知识图谱graphiti)
  - [集成 GitHub 与 Google OAuth](#集成-github-与-google-oauth)
  - [Docker 镜像配置](#docker-镜像配置)
- [开发](#开发)
- [测试 LLM 智能体](#测试-llm-智能体)
- [嵌入模型配置与测试](#嵌入模型配置与测试)
- [使用 ftester 测试函数](#使用-ftester-测试函数)
- [构建](#构建)
- [致谢](#致谢)
- [许可证](#许可证)

## 概述

PentAGI 是一款用于自动化安全测试的创新工具，采用了前沿的人工智能技术。本项目面向需要强大且灵活的渗透测试方案的信息安全从业者、研究人员和爱好者。

你可以观看视频 **PentAGI 概览**：
[![PentAGI 概览视频](https://github.com/user-attachments/assets/0828dc3e-15f1-4a1d-858e-9696a146e478)](https://youtu.be/R70x5Ddzs1o)

## 功能特性

- 安全隔离。所有操作都在完全隔离的 Docker 沙箱环境中执行。
- 完全自主。AI 智能体可自主判断并执行渗透测试步骤，还可按需启用执行监控和智能任务规划，提高执行可靠性。
- 专业渗透测试工具。内置 20 多款专业安全工具，包括 nmap、metasploit、sqlmap 等。
- 智能记忆系统。长期保存研究成果和成功方案，供后续复用。
- 知识图谱集成。基于 Graphiti 和 Neo4j 构建知识图谱，用于语义关系追踪和深层上下文理解。
- 网络情报采集。通过内置浏览器 [scraper](https://hub.docker.com/r/vxcontrol/scraper) 从网络来源获取最新信息。
- 外部搜索系统。集成多种高级搜索 API，包括 [Tavily](https://tavily.com)、[Traversaal](https://traversaal.ai)、[Perplexity](https://www.perplexity.ai)、[DuckDuckGo](https://duckduckgo.com/)、[Google Custom Search](https://programmablesearchengine.google.com/)、[Sploitus Search](https://sploitus.com) 和 [Searxng](https://searxng.org)，可更全面地采集信息。
- 专家团队协作。任务委派系统会让专职 AI 智能体分别负责研究、开发和基础设施任务；配合可选的执行监控与智能任务规划，可以让小型模型取得更理想的执行效果。
- 全面监控。提供详尽日志，并集成 Grafana/Prometheus，便于实时掌握系统状态。
- 详细报告。生成详尽的漏洞报告，并附带漏洞利用指南。
- 智能容器管理。根据具体任务需求自动选择 Docker 镜像。
- 现代化界面。简洁直观的网页界面，用于系统管理与监控。
- 完备的 API。功能齐全的 REST 与 GraphQL API，支持 Bearer 令牌认证，便于自动化与系统集成。
- 持久化存储。所有命令与输出均存储在带 [pgvector](https://hub.docker.com/r/vxcontrol/pgvector) 扩展的 PostgreSQL 中。
- 可扩展架构。基于微服务设计，支持横向扩展。
- 自托管方案。完全掌控自己的部署环境与数据。
- 灵活的模型接入。支持 10 多个 LLM 提供商（[OpenAI](https://platform.openai.com/)、[Anthropic](https://www.anthropic.com/)、[Google AI/Gemini](https://ai.google.dev/)、[AWS Bedrock](https://aws.amazon.com/bedrock/)、[Ollama](https://ollama.com/)、[DeepSeek](https://www.deepseek.com/en/)、[GLM](https://z.ai/)、[Kimi](https://platform.moonshot.ai/)、[Qwen](https://www.alibabacloud.com/en/) 及自定义提供商），也支持聚合服务（[OpenRouter](https://openrouter.ai/)、[DeepInfra](https://deepinfra.com/)）。如需在生产环境中本地部署，请参阅 [vLLM + Qwen3.5-27B-FP8 指南](examples/guides/vllm-qwen35-27b-fp8.md)。
- API 令牌认证。通过安全的 Bearer 令牌机制，供程序调用 REST 与 GraphQL API。
- 快速部署。通过 [Docker Compose](https://docs.docker.com/compose/) 轻松搭建，并提供完善的环境变量配置。

### 当前能力边界

- PentAGI 目前是一款可自主运行、也可由助手引导的渗透测试平台，并非 CALDERA 那类带有预设攻击活动或攻击计划的入侵与攻击模拟（BAS）或对手模拟产品。
- 由智能体自行编写的 BAS 式攻击脚本应视为设想或未来工作，而非现已实现的功能。
- 当前的任务流报告界面支持网页查看、复制到剪贴板，以及下载 Markdown 或 PDF；JSON 格式尚未列入支持的任务流报告导出格式。
- 目前既可使用内置提供商，也可通过自定义或 OpenAI 兼容端点灵活接入模型。参见 [自定义 LLM 提供商配置](#自定义-llm-提供商配置)和 [vLLM + Qwen3.5-27B-FP8 指南](examples/guides/vllm-qwen35-27b-fp8.md)。

## 架构

### 系统上下文

```mermaid
flowchart TB
    classDef person fill:#08427B,stroke:#073B6F,color:#fff
    classDef system fill:#1168BD,stroke:#0B4884,color:#fff
    classDef external fill:#666666,stroke:#0B4884,color:#fff

    pentester["👤 Security Engineer
    (User of the system)"]

    pentagi["✨ PentAGI
    (Autonomous penetration testing system)"]

    target["🎯 target-system
    (System under test)"]
    llm["🧠 llm-provider
    (OpenAI/Anthropic/Ollama/Bedrock/Gemini/Custom)"]
    search["🔍 search-systems
    (Google/DuckDuckGo/Tavily/Traversaal/Perplexity/Sploitus/Searxng)"]
    langfuse["📊 langfuse-ui
    (LLM Observability Dashboard)"]
    grafana["📈 grafana
    (System Monitoring Dashboard)"]

    pentester --> |Uses HTTPS| pentagi
    pentester --> |Monitors AI HTTPS| langfuse
    pentester --> |Monitors System HTTPS| grafana
    pentagi --> |Tests Various protocols| target
    pentagi --> |Queries HTTPS| llm
    pentagi --> |Searches HTTPS| search
    pentagi --> |Reports HTTPS| langfuse
    pentagi --> |Reports HTTPS| grafana

    class pentester person
    class pentagi system
    class target,llm,search,langfuse,grafana external

    linkStyle default stroke:#ffffff,color:#ffffff
```

<details>
<summary><b>容器架构</b>（点击展开）</summary>

```mermaid
graph TB
    subgraph Core Services
        UI[Frontend UI<br/>React + TypeScript]
        API[Backend API<br/>Go + GraphQL]
        DB[(Vector Store<br/>PostgreSQL + pgvector)]
        MQ[Task Queue<br/>Async Processing]
        Agent[AI Agents<br/>Multi-Agent System]
    end

    subgraph Knowledge Graph
        Graphiti[Graphiti<br/>Knowledge Graph API]
        Neo4j[(Neo4j<br/>Graph Database)]
    end

    subgraph Monitoring
        Grafana[Grafana<br/>Dashboards]
        VictoriaMetrics[VictoriaMetrics<br/>Time-series DB]
        Jaeger[Jaeger<br/>Distributed Tracing]
        Loki[Loki<br/>Log Aggregation]
        OTEL[OpenTelemetry<br/>Data Collection]
    end

    subgraph Analytics
        Langfuse[Langfuse<br/>LLM Analytics]
        ClickHouse[ClickHouse<br/>Analytics DB]
        Redis[Redis<br/>Cache + Rate Limiter]
        MinIO[MinIO<br/>S3 Storage]
    end

    subgraph Security Tools
        Scraper[Web Scraper<br/>Isolated Browser]
        PenTest[Security Tools<br/>20+ Pro Tools<br/>Sandboxed Execution]
    end

    UI --> |HTTP/WS| API
    API --> |SQL| DB
    API --> |Events| MQ
    MQ --> |Tasks| Agent
    Agent --> |Commands| PenTest
    Agent --> |Queries| DB
    Agent --> |Knowledge| Graphiti
    Graphiti --> |Graph| Neo4j

    API --> |Telemetry| OTEL
    OTEL --> |Metrics| VictoriaMetrics
    OTEL --> |Traces| Jaeger
    OTEL --> |Logs| Loki

    Grafana --> |Query| VictoriaMetrics
    Grafana --> |Query| Jaeger
    Grafana --> |Query| Loki

    API --> |Analytics| Langfuse
    Langfuse --> |Store| ClickHouse
    Langfuse --> |Cache| Redis
    Langfuse --> |Files| MinIO

    classDef core fill:#f9f,stroke:#333,stroke-width:2px,color:#000
    classDef knowledge fill:#ffa,stroke:#333,stroke-width:2px,color:#000
    classDef monitoring fill:#bbf,stroke:#333,stroke-width:2px,color:#000
    classDef analytics fill:#bfb,stroke:#333,stroke-width:2px,color:#000
    classDef tools fill:#fbb,stroke:#333,stroke-width:2px,color:#000

    class UI,API,DB,MQ,Agent core
    class Graphiti,Neo4j knowledge
    class Grafana,VictoriaMetrics,Jaeger,Loki,OTEL monitoring
    class Langfuse,ClickHouse,Redis,MinIO analytics
    class Scraper,PenTest tools
```

</details>

<details>
<summary><b>实体关系</b>（点击展开）</summary>

```mermaid
erDiagram
    Flow ||--o{ Task : contains
    Task ||--o{ SubTask : contains
    SubTask ||--o{ Action : contains
    Action ||--o{ Artifact : produces
    Action ||--o{ Memory : stores

    Flow {
        string id PK
        string name "Flow name"
        string description "Flow description"
        string status "active/completed/failed"
        json parameters "Flow parameters"
        timestamp created_at
        timestamp updated_at
    }

    Task {
        string id PK
        string flow_id FK
        string name "Task name"
        string description "Task description"
        string status "pending/running/done/failed"
        json result "Task results"
        timestamp created_at
        timestamp updated_at
    }

    SubTask {
        string id PK
        string task_id FK
        string name "Subtask name"
        string description "Subtask description"
        string status "queued/running/completed/failed"
        string agent_type "researcher/developer/executor"
        json context "Agent context"
        timestamp created_at
        timestamp updated_at
    }

    Action {
        string id PK
        string subtask_id FK
        string type "command/search/analyze/etc"
        string status "success/failure"
        json parameters "Action parameters"
        json result "Action results"
        timestamp created_at
    }

    Artifact {
        string id PK
        string action_id FK
        string type "file/report/log"
        string path "Storage path"
        json metadata "Additional info"
        timestamp created_at
    }

    Memory {
        string id PK
        string action_id FK
        string type "observation/conclusion"
        vector embedding "Vector representation"
        text content "Memory content"
        timestamp created_at
    }
```

</details>

<details>
<summary><b>智能体交互</b>（点击展开）</summary>

```mermaid
sequenceDiagram
    participant O as Orchestrator
    participant R as Researcher
    participant D as Developer
    participant E as Executor
    participant VS as Vector Store
    participant KB as Knowledge Base

    Note over O,KB: Flow Initialization
    O->>VS: Query similar tasks
    VS-->>O: Return experiences
    O->>KB: Load relevant knowledge
    KB-->>O: Return context

    Note over O,R: Research Phase
    O->>R: Analyze target
    R->>VS: Search similar cases
    VS-->>R: Return patterns
    R->>KB: Query vulnerabilities
    KB-->>R: Return known issues
    R->>VS: Store findings
    R-->>O: Research results

    Note over O,D: Planning Phase
    O->>D: Plan attack
    D->>VS: Query exploits
    VS-->>D: Return techniques
    D->>KB: Load tools info
    KB-->>D: Return capabilities
    D-->>O: Attack plan

    Note over O,E: Execution Phase
    O->>E: Execute plan
    E->>KB: Load tool guides
    KB-->>E: Return procedures
    E->>VS: Store results
    E-->>O: Execution status
```

</details>

<details>
<summary><b>记忆系统</b>（点击展开）</summary>

```mermaid
graph TB
    subgraph "Long-term Memory"
        VS[(Vector Store<br/>Embeddings DB)]
        KB[Knowledge Base<br/>Domain Expertise]
        Tools[Tools Knowledge<br/>Usage Patterns]
    end

    subgraph "Working Memory"
        Context[Current Context<br/>Task State]
        Goals[Active Goals<br/>Objectives]
        State[System State<br/>Resources]
    end

    subgraph "Episodic Memory"
        Actions[Past Actions<br/>Commands History]
        Results[Action Results<br/>Outcomes]
        Patterns[Success Patterns<br/>Best Practices]
    end

    Context --> |Query| VS
    VS --> |Retrieve| Context

    Goals --> |Consult| KB
    KB --> |Guide| Goals

    State --> |Record| Actions
    Actions --> |Learn| Patterns
    Patterns --> |Store| VS

    Tools --> |Inform| State
    Results --> |Update| Tools

    VS --> |Enhance| KB
    KB --> |Index| VS

    classDef ltm fill:#f9f,stroke:#333,stroke-width:2px,color:#000
    classDef wm fill:#bbf,stroke:#333,stroke-width:2px,color:#000
    classDef em fill:#bfb,stroke:#333,stroke-width:2px,color:#000

    class VS,KB,Tools ltm
    class Context,Goals,State wm
    class Actions,Results,Patterns em
```

</details>

<details>
<summary><b>对话链摘要</b>（点击展开）</summary>

对话链摘要系统会有选择地概括较早的消息，避免上下文不断增长。这样既不会超出 token 上限，也能保持对话连贯。

```mermaid
flowchart TD
    A[Input Chain] --> B{Needs Summarization?}
    B -->|No| C[Return Original Chain]
    B -->|Yes| D[Convert to ChainAST]
    D --> E[Apply Section Summarization]
    E --> F[Process Oversized Pairs]
    F --> G[Manage Last Section Size]
    G --> H[Apply QA Summarization]
    H --> I[Rebuild Chain with Summaries]
    I --> J{Is New Chain Smaller?}
    J -->|Yes| K[Return Optimized Chain]
    J -->|No| C

    classDef process fill:#bbf,stroke:#333,stroke-width:2px,color:#000
    classDef decision fill:#bfb,stroke:#333,stroke-width:2px,color:#000
    classDef output fill:#fbb,stroke:#333,stroke-width:2px,color:#000

    class A,D,E,F,G,H,I process
    class B,J decision
    class C,K output
```

该算法使用对话链的结构化表示（ChainAST），保留各类消息，包括工具调用及其响应。所有摘要操作都会在缩减上下文的同时，保持关键对话流程完整。

### 全局摘要器配置项

| 参数             | 环境变量                         | 默认值  | 说明                                     |
| ---------------- | -------------------------------- | ------- | ---------------------------------------- |
| 保留最后一节     | `SUMMARIZER_PRESERVE_LAST`       | `true`  | 是否完整保留最后一节中的所有消息         |
| 使用问答对       | `SUMMARIZER_USE_QA`              | `true`  | 是否采用问答对摘要策略                   |
| 摘要问答中的用户消息 | `SUMMARIZER_SUM_MSG_HUMAN_IN_QA` | `false` | 是否对问答对中的用户消息进行摘要         |
| 最后一节大小     | `SUMMARIZER_LAST_SEC_BYTES`      | `51200` | 最后一节的最大字节数（50KB）             |
| 单个消息对大小上限 | `SUMMARIZER_MAX_BP_BYTES`        | `16384` | 单个消息对的最大字节数（16KB）           |
| 最大问答节数     | `SUMMARIZER_MAX_QA_SECTIONS`     | `10`    | 最多保留的问答对节数                     |
| 问答内容上限     | `SUMMARIZER_MAX_QA_BYTES`        | `65536` | 问答对各节的最大总字节数（64KB）         |
| 保留近期问答节数 | `SUMMARIZER_KEEP_QA_SECTIONS`    | `1`     | 保留多少个近期问答节而不进行摘要         |

### 助手摘要器配置项

助手实例可以使用独立的摘要设置，对上下文管理策略进行细粒度调整：

| 参数             | 环境变量                                | 默认值  | 说明                                        |
| ---------------- | --------------------------------------- | ------- | ------------------------------------------- |
| 保留最后一节     | `ASSISTANT_SUMMARIZER_PRESERVE_LAST`    | `true`  | 是否完整保留助手最后一节中的所有消息        |
| 最后一节大小     | `ASSISTANT_SUMMARIZER_LAST_SEC_BYTES`   | `76800` | 助手最后一节的最大字节数（75KB）            |
| 单个消息对大小上限 | `ASSISTANT_SUMMARIZER_MAX_BP_BYTES`     | `16384` | 助手上下文中单个消息对的最大字节数（16KB）  |
| 最大问答节数     | `ASSISTANT_SUMMARIZER_MAX_QA_SECTIONS`  | `7`     | 助手上下文中最多保留的问答节数              |
| 问答内容上限     | `ASSISTANT_SUMMARIZER_MAX_QA_BYTES`     | `76800` | 助手问答内容的最大总字节数（75KB）          |
| 保留近期问答节数 | `ASSISTANT_SUMMARIZER_KEEP_QA_SECTIONS` | `3`     | 保留多少个近期问答节而不进行摘要            |

与全局设置相比，助手摘要器会保留更大的上下文容量，在保存更多近期对话的同时兼顾 token 使用效率。

### 摘要器环境变量配置

```bash
# 全局摘要逻辑的默认值
SUMMARIZER_PRESERVE_LAST=true
SUMMARIZER_USE_QA=true
SUMMARIZER_SUM_MSG_HUMAN_IN_QA=false
SUMMARIZER_LAST_SEC_BYTES=51200
SUMMARIZER_MAX_BP_BYTES=16384
SUMMARIZER_MAX_QA_SECTIONS=10
SUMMARIZER_MAX_QA_BYTES=65536
SUMMARIZER_KEEP_QA_SECTIONS=1

# 助手摘要逻辑的默认值
ASSISTANT_SUMMARIZER_PRESERVE_LAST=true
ASSISTANT_SUMMARIZER_LAST_SEC_BYTES=76800
ASSISTANT_SUMMARIZER_MAX_BP_BYTES=16384
ASSISTANT_SUMMARIZER_MAX_QA_SECTIONS=7
ASSISTANT_SUMMARIZER_MAX_QA_BYTES=76800
ASSISTANT_SUMMARIZER_KEEP_QA_SECTIONS=3
```

</details>

<a id="advanced-agent-supervision"></a>
<details>
<summary><b>高级智能体监管</b>（点击展开）</summary>

PentAGI 内置多层智能体监管机制，能够提高任务执行效率、防止无限循环，并在执行陷入停滞时智能恢复：

### 执行监控（测试版）
- **导师自动介入**：检测到可能存在异常的执行模式时，自动启用顾问智能体（导师）
- **模式检测**：监测同一工具的重复调用（默认阈值 5，可配置）和工具调用总次数（默认阈值 10，可配置）
- **进展分析**：评估智能体是否在推进子任务，识别循环和低效行为
- **替代策略**：当前策略失效时推荐其他思路
- **信息检索引导**：建议先查找已有解决方案，避免重复造轮子
- **扩展响应格式**：工具响应中同时包含 `<original_result>` 与 `<mentor_analysis>` 两部分
- **可配置**：通过 `EXECUTION_MONITOR_ENABLED` 启用（默认关闭），并可用 `EXECUTION_MONITOR_SAME_TOOL_LIMIT` 和 `EXECUTION_MONITOR_TOTAL_TOOL_LIMIT` 自定义阈值

**适用场景**：适合参数量较小的模型（< 32B）和需要持续引导的复杂攻击场景，也能避免智能体反复尝试同一种方案。

**性能影响**：执行时间与 token 用量会增至原来的 2-3 倍；但 Qwen3.5-27B-FP8 的测试表明，**结果质量约为基线的 2 倍**。

### 智能任务规划（测试版）
- **自动拆解**：在专职智能体开始工作前，由 Planner（规划器，即规划模式下的 Adviser）生成 3-7 个具体且可执行的步骤
- **上下文感知规划**：通过 Enricher（信息增强智能体）分析完整的执行上下文，据此制定计划
- **结构化分配**：将原始请求封装进 `<task_assignment>` 结构，并附带执行计划与说明
- **范围管理**：让智能体仅专注于当前子任务，防止任务范围不断扩张
- **补充执行要点**：计划会标出关键操作、潜在问题和验证点
- **可配置**：通过 `AGENT_PLANNING_STEP_ENABLED` 启用（默认关闭）

**适用场景**：参数量 < 32B 的模型、复杂的渗透测试工作流，以及需要提高复杂任务成功率的场景

**强化顾问配置**：顾问智能体使用更强的模型或增强设置时，效果尤为明显。例如，让顾问使用同一个基础模型并开启最高推理模式（参见 [`vllm-qwen3.5-27b-fp8.provider.yml`](examples/configs/vllm-qwen3.5-27b-fp8.provider.yml)），无需更换模型架构即可完成全面的任务分析和策略规划。

**性能影响**：会增加规划开销，但能显著提高任务完成率并减少重复工作

### 工具调用上限（始终生效）
- **硬性上限**：无论监管功能是否启用，都能防止任务失控执行
- **按智能体类型区分**：
  - 通用智能体（Assistant、Primary Agent、Pentester、Coder、Installer）：`MAX_GENERAL_AGENT_TOOL_CALLS`（默认 100）
  - 受限智能体（Searcher、Enricher、Memorist、Generator、Reporter、Adviser、Reflector、Planner）：`MAX_LIMITED_AGENT_TOOL_CALLS`（默认 20）
- **平稳终止**：接近上限时，由 Reflector 引导智能体妥善结束任务
- **资源保护**：保障系统稳定，防止资源耗尽

### Reflector 集成（始终生效）
- **自动纠正**：LLM 连续 3 次未能生成工具调用时自动触发
- **策略指导**：分析失败原因，引导智能体正确调用工具，或调用 `done`、`ask` 等 barrier 工具来结束或暂停流程
- **恢复机制**：根据具体的失败模式，结合上下文给出有针对性的指导
- **上限控制**：达到工具调用上限时，协调智能体平稳结束任务

### 面向开源模型的建议

**参数量 < 32B 的模型必备**：
Qwen3.5-27B-FP8 的测试表明，对参数量较小的开源模型而言，同时启用执行监控与任务规划是取得稳定结果的必要条件：
- **质量提升**：与未启用监管的基线执行相比，结果质量约为基线的 2 倍
- **防止循环**：显著减少无限循环和重复工作
- **攻击多样性**：鼓励探索多种攻击向量，避免固守单一思路
- **网络隔离部署**：配合本地 LLM 推理，可在与外网隔离的环境中实现生产级自主渗透测试

**权衡取舍**：
- token 用量：由于需要调用导师和规划器，会增至原来的 2-3 倍
- 执行时间：由于增加了分析与规划步骤，会延长至原来的 2-3 倍
- 结果质量：完整性、准确性和攻击覆盖范围约为基线的 2 倍
- 模型要求：顾问采用增强配置时效果最佳，例如提高推理强度、使用更强的模型版本或更换模型

**配置策略**：
要让小型模型取得更好的效果，可为顾问智能体使用增强配置：
- 使用同一个模型并开启最高推理模式（示例：[`vllm-qwen3.5-27b-fp8.provider.yml`](examples/configs/vllm-qwen3.5-27b-fp8.provider.yml)）
- 或让顾问使用更强的模型，其他智能体继续使用基础模型
- 根据任务复杂度和模型能力调整监控阈值



</details>

PentAGI 的架构遵循模块化、可扩展和安全的设计原则，主要组件如下：

1. **核心服务**
   - 前端 UI：基于 React 构建的网页界面，由 TypeScript 提供类型安全保障
   - 后端 API：基于 Go 构建的 REST 和 GraphQL API，支持 Bearer 令牌认证，供程序调用
   - 向量存储：采用启用了 pgvector 扩展的 PostgreSQL，实现语义搜索和记忆存储
   - 任务队列：可靠处理异步任务
   - AI 智能体：由多个专职智能体组成，可高效执行测试

2. **知识图谱**
   - Graphiti：用于追踪语义关系和理解上下文的知识图谱 API
   - Neo4j：用于存储和查询实体、操作与结果之间关系的图数据库
   - 自动记录智能体响应和工具执行过程，逐步形成完整的知识库

3. **监控技术栈**
   - OpenTelemetry：统一采集并关联可观测性数据
   - Grafana：提供实时可视化与告警仪表板
   - VictoriaMetrics：高性能时序指标存储
   - Jaeger：用于调试的端到端分布式追踪
   - Loki：可扩展的日志聚合与分析系统

4. **分析平台**
   - Langfuse：提供 LLM 深度可观测性和性能分析
   - ClickHouse：列式分析数据仓库
   - Redis：用于高速缓存和速率限制
   - MinIO：用于存储产物的 S3 兼容对象存储

5. **安全工具**
   - 网页抓取器：用于安全访问网页的隔离浏览器环境
   - 渗透测试工具：覆盖 20 多种专业安全工具
   - 沙箱执行：所有操作均在隔离容器中运行

6. **记忆系统**
   - 长期记忆：持久保存知识与经验
   - 工作记忆：保存当前操作所需的上下文与目标
   - 情景记忆：记录历史操作与成功模式
   - 知识库：存储结构化的领域知识与工具能力
   - 上下文管理：通过对话链摘要智能管理不断增长的 LLM 上下文窗口

系统通过 Docker 容器实现隔离并简化部署。核心服务、监控组件和分析组件分别使用独立网络，形成清晰的安全边界。各组件均支持横向扩展，也可在生产环境中配置为高可用。

## 快速开始

### 系统要求

- Docker 与 Docker Compose（或 Podman，参见 [Podman 配置](#使用-podman-运行-pentagi)）
- 至少 2 个 vCPU
- 至少 4GB 内存
- 20GB 可用磁盘空间
- 可访问互联网，用于下载镜像和更新

### 使用汉化镜像（推荐）

```bash
git clone --branch zh-CN https://github.com/fulaoaz/pentagi.git
cd pentagi
cp .env.example .env
```

首次启动前，请至少修改 `.env` 中的 `COOKIE_SIGNING_SALT` 和 `PENTAGI_POSTGRES_PASSWORD`，并按需填写一个 LLM 提供商的密钥或本地模型地址。`PENTAGI_IMAGE` 已默认指向 `ghcr.io/fulaoaz/pentagi:zh-cn`。

```bash
docker compose pull
docker compose up -d
docker compose ps
```

打开 [https://localhost:8443](https://localhost:8443)，使用默认账号 `admin@pentagi.com` / `admin` 登录，并按首次登录提示修改密码。默认使用本机自签名证书，浏览器首次访问时需要确认该本地证书。

以后更新汉化版只需在仓库目录执行：

```bash
git pull --ff-only
docker compose pull
docker compose up -d
```

### 使用汉化安装程序

本仓库自带 Go 编写的终端交互式安装器，源码位于 [`backend/cmd/installer`](backend/cmd/installer)。安装器界面、系统检查、配置向导、部署进度、维护菜单和密码重置提示均已汉化；命令、环境变量、提供商 ID 和原始技术错误仍保留英文，便于排障和对照上游文档。

每次 `zh-CN` 分支更新后，GitHub Actions 都会重新构建并覆盖发布 [`zh-cn-installer-latest`](https://github.com/fulaoaz/pentagi/releases/tag/zh-cn-installer-latest)。安装器内嵌当前汉化分支的 Compose、`.env` 模板和示例配置，默认部署 `ghcr.io/fulaoaz/pentagi:zh-cn`。

**支持的平台：**
- **Linux**：amd64 [下载](https://github.com/fulaoaz/pentagi/releases/download/zh-cn-installer-latest/pentagi-installer-zh-cn-linux-amd64.zip) | arm64 [下载](https://github.com/fulaoaz/pentagi/releases/download/zh-cn-installer-latest/pentagi-installer-zh-cn-linux-arm64.zip)
- **Windows**：amd64 [下载](https://github.com/fulaoaz/pentagi/releases/download/zh-cn-installer-latest/pentagi-installer-zh-cn-windows-amd64.zip)
- **macOS**：amd64（Intel）[下载](https://github.com/fulaoaz/pentagi/releases/download/zh-cn-installer-latest/pentagi-installer-zh-cn-darwin-amd64.zip) | arm64（Apple Silicon）[下载](https://github.com/fulaoaz/pentagi/releases/download/zh-cn-installer-latest/pentagi-installer-zh-cn-darwin-arm64.zip)
- **完整性校验**：[SHA256SUMS](https://github.com/fulaoaz/pentagi/releases/download/zh-cn-installer-latest/SHA256SUMS)

**快速安装（Linux amd64）：**

```bash
# 创建安装目录
mkdir -p pentagi && cd pentagi

# 下载汉化安装程序
wget -O installer.zip https://github.com/fulaoaz/pentagi/releases/download/zh-cn-installer-latest/pentagi-installer-zh-cn-linux-amd64.zip

# 解压
unzip installer.zip

# 运行交互式安装程序
chmod +x installer
sudo ./installer
```

Windows 用户需先启动 Docker Desktop，解压后在 PowerShell 中运行：

```powershell
.\installer.exe
```

也可以在 Linux、WSL 或 macOS 中直接从当前分支源码构建。以下命令会先把 Compose、`.env` 模板和示例配置嵌入安装器，再生成可独立运行的二进制文件；Windows 用户建议直接下载上面的 `.zip` 成品：

```bash
git clone --branch zh-CN https://github.com/fulaoaz/pentagi.git
cd pentagi/backend
go generate ./cmd/installer/files
go build -trimpath -o ../installer ./cmd/installer
cd ..
./installer
```

仓库中容易混淆的安装相关文件如下：

| 路径 | 用途 | 是否需要汉化 |
| --- | --- | --- |
| `backend/cmd/installer` | 面向用户的交互式安装、配置和维护程序 | 已汉化 |
| `docker-compose*.yml`、`.env.example` | 安装器和手动部署共用的配置模板 | 中文说明已补充，变量名保持英文 |
| `scripts/version.sh`、`scripts/version.ps1` | 开发和发布时计算版本号 | 开发脚本，无需汉化 |
| `scripts/entrypoint.sh` | 容器内部启动入口 | 内部脚本，无需汉化 |

> `pentagi.com/downloads/.../installer-latest.zip` 是上游官方安装器，不是此分支构建的汉化安装器。要确保安装界面和默认镜像都来自汉化分支，请使用上面的 GitHub Release 下载地址。

**前置条件与权限：**

安装程序需要具备相应权限才能与 Docker API 交互并正常工作。默认情况下它使用 Docker 套接字（`/var/run/docker.sock`），因此需要满足以下任一条件：

- **方式 1（生产环境推荐）：** 以 root 身份运行安装程序：
  ```bash
  sudo ./installer
  ```

- **方式 2（开发环境）：** 将你的用户加入 `docker` 组，以获得 Docker 套接字的访问权限：
  ```bash
  # 将当前用户加入 docker 组
  sudo usermod -aG docker $USER
  
  # 注销后重新登录，或立即激活该组
  newgrp docker
  
  # 验证 Docker 访问权限（应无需 sudo 即可执行）
  docker ps
  ```

  ⚠️ **安全提示：** 将用户加入 `docker` 组等同于授予 root 权限。请仅在受控环境中对可信用户这样做。生产部署建议使用 Docker 无根模式，或通过 sudo 运行安装程序。

安装程序会执行以下步骤：
1. **系统检查**：验证 Docker、网络连通性和系统要求
2. **环境准备**：创建 `.env` 文件并写入合适的默认配置
3. **提供商配置**：设置 LLM 提供商（OpenAI、Anthropic、Gemini、Bedrock、Ollama、自定义）
4. **搜索引擎**：配置 DuckDuckGo、Google、Tavily、Traversaal、Perplexity、Sploitus、Searxng
5. **安全加固**：生成安全凭据并配置 SSL 证书
6. **部署**：通过 docker-compose 启动 PentAGI

### 当前网页端可配置的内容

服务器启动后，可直接在 PentAGI 网页控制台管理以下设置：

- **Settings -> Providers（设置 -> 提供商）**：为受支持的提供商类型创建、编辑、删除和测试用户自定义配置。每项配置可指定各智能体使用的模型、运行参数、推理选项和价格元数据。
- **Settings -> Prompts（设置 -> 提示词）**：管理系统、用户和工具提示词模板。
- **Settings -> PentAGI API（设置 -> PentAGI API）**：创建和管理用于访问 REST 与 GraphQL API 的 PentAGI Bearer 令牌。
- **其他由界面管理的偏好**：收藏的任务流会保存到用户偏好中；主题则从主侧边栏的个人资料菜单中选择，不在设置页面中配置。

### 仍需在服务端配置的内容

以下项目仍需通过环境变量、Compose 文件或挂载的配置文件在服务端设置：

- **LLM 凭据和连接信息**：OpenAI、Anthropic、Bedrock、Ollama、自定义提供商及类似后端所需的 API 密钥、端点、认证方式和提供商专用连接设置。仅部分提供商支持通过配置文件路径加载设置，例如 `OLLAMA_SERVER_CONFIG_PATH` 和 `LLM_SERVER_CONFIG_PATH`。
- **搜索提供商凭据和选项**：包括 `DUCKDUCKGO_*`、`GOOGLE_*`、`TAVILY_API_KEY`、`TRAVERSAAL_API_KEY`、`PERPLEXITY_*`、`SEARXNG_*` 和 `SPLOITUS_ENABLED` 等设置。
- **第三方集成**：Langfuse、Graphiti 及类似外部服务仍需在服务端配置。
- **MCP 服务器管理**：网页控制台目前尚未提供可直接使用的 MCP 设置页面。

**生产环境与更高安全要求：**

对于生产部署或安全敏感环境，**强烈建议**采用分布式双节点架构，将 Worker 操作隔离到另一台服务器。这样可避免不受信任的代码执行和网络访问给主系统带来风险。

**详细指南**：[Worker 节点配置](examples/guides/worker_node.md)

双节点方案具有以下特点：
- **隔离执行**：Worker 容器运行在专用硬件上
- **网络隔离**：为渗透测试划分独立的网络边界
- **安全边界**：通过 TLS 认证保护 Docker-in-Docker
- **OOB（带外）攻击支持**：为带外技术预留独立的端口范围

### 手动安装

1. 克隆汉化仓库并进入目录：

```bash
git clone --branch zh-CN https://github.com/fulaoaz/pentagi.git
cd pentagi
```

2. 将 `.env.example` 复制为 `.env`；如果只下载部署文件，请从汉化分支获取：

```bash
cp .env.example .env
# 或：curl -o .env https://raw.githubusercontent.com/fulaoaz/pentagi/zh-CN/.env.example
```

3. 创建示例配置文件（`example.custom.provider.yml`、`example.ollama.provider.yml`），或下载现成配置：

```bash
curl -o example.custom.provider.yml https://raw.githubusercontent.com/fulaoaz/pentagi/zh-CN/examples/configs/custom-openai.provider.yml
curl -o example.ollama.provider.yml https://raw.githubusercontent.com/fulaoaz/pentagi/zh-CN/examples/configs/ollama-llama318b.provider.yml
```

4. 在 `.env` 文件中填写所需的 API 密钥。

```bash
# 必填：至少选择下列一个 LLM 提供商
OPEN_AI_KEY=your_openai_key
ANTHROPIC_API_KEY=your_anthropic_key
GEMINI_API_KEY=your_gemini_key

# 可选：AWS Bedrock 提供商（企业级模型）
BEDROCK_REGION=us-east-1
# 选择一种认证方式：
BEDROCK_DEFAULT_AUTH=true                        # 选项 1：使用 AWS SDK 默认凭据链（建议用于 EC2/ECS）
# BEDROCK_BEARER_TOKEN=your_bearer_token         # 选项 2：Bearer 令牌认证
# BEDROCK_ACCESS_KEY_ID=your_aws_access_key      # 选项 3：静态凭据
# BEDROCK_SECRET_ACCESS_KEY=your_aws_secret_key

# 可选：Ollama 提供商（本地或云端）
# OLLAMA_SERVER_URL=http://ollama-server:11434   # 本地服务器
# OLLAMA_SERVER_URL=https://ollama.com           # 云服务
# OLLAMA_SERVER_API_KEY=your_ollama_cloud_key    # 云服务必填，本地服务留空

# 可选：中国 AI 提供商
# DEEPSEEK_API_KEY=your_deepseek_key             # DeepSeek（推理能力强）
# GLM_API_KEY=your_glm_key                       # GLM（智谱 AI）
# KIMI_API_KEY=your_kimi_key                     # Kimi（Moonshot AI，超长上下文）
# QWEN_API_KEY=your_qwen_key                     # Qwen（阿里云，多模态）

# 可选：本地 LLM 提供商（推理不产生 API 费用）
OLLAMA_SERVER_URL=http://localhost:11434
OLLAMA_SERVER_MODEL=your_model_name

# 可选：其他搜索能力
DUCKDUCKGO_ENABLED=true
DUCKDUCKGO_REGION=us-en
DUCKDUCKGO_SAFESEARCH=
DUCKDUCKGO_TIME_RANGE=
SPLOITUS_ENABLED=true
GOOGLE_API_KEY=your_google_key
GOOGLE_CX_KEY=your_google_cx
TAVILY_API_KEY=your_tavily_key
TRAVERSAAL_API_KEY=your_traversaal_key
PERPLEXITY_API_KEY=your_perplexity_key
PERPLEXITY_MODEL=sonar-pro
PERPLEXITY_CONTEXT_SIZE=medium

# Searxng 元搜索引擎（聚合多个来源的结果）
SEARXNG_URL=http://your-searxng-instance:8080
SEARXNG_CATEGORIES=general
SEARXNG_LANGUAGE=
SEARXNG_SAFESEARCH=0
SEARXNG_TIME_RANGE=
SEARXNG_TIMEOUT=

## Graphiti 知识图谱设置
GRAPHITI_ENABLED=true
GRAPHITI_TIMEOUT=30
GRAPHITI_URL=http://graphiti:8000
GRAPHITI_MODEL_NAME=gpt-5-mini

# Neo4j 设置（供 Graphiti 服务使用）
NEO4J_USER=neo4j
NEO4J_DATABASE=neo4j
NEO4J_PASSWORD=devpassword
NEO4J_URI=bolt://neo4j:7687

# 助手配置
ASSISTANT_USE_AGENTS=false         # 新建助手时“使用智能体”的默认值
```

5. 修改 `.env` 文件中所有与安全相关的环境变量，加强系统安全性。

<details>
    <summary>与安全相关的环境变量</summary>

### 主要安全设置
- `COOKIE_SIGNING_SALT` —— 用于 Cookie 签名的盐值，请改为随机值
- `PUBLIC_URL` —— 服务器的公开 URL（例如 `https://pentagi.example.com`）
- `SERVER_SSL_CRT` 和 `SERVER_SSL_KEY` —— 已有 HTTPS 证书和私钥的路径（需在 docker-compose.yml 中以卷的形式挂载）

### 抓取服务访问
- `SCRAPER_PUBLIC_URL` —— 如需使用另一台抓取服务器访问公网 URL，请在此填写其公开地址
- `SCRAPER_PRIVATE_URL` —— 抓取服务的内部地址（即 docker-compose.yml 中用于访问本地地址的抓取服务）

### 访问凭据
- `PENTAGI_POSTGRES_USER` 和 `PENTAGI_POSTGRES_PASSWORD` —— PostgreSQL 凭据
- `NEO4J_USER` 和 `NEO4J_PASSWORD` —— Neo4j 凭据（用于 Graphiti 知识图谱）

</details>

6. 如果希望在 VSCode 或其他 IDE 中把 `.env` 作为 envFile 使用，可移除文件中所有行内注释：

```bash
perl -i -pe 's/\s+#.*$//' .env
```

7. 启动 PentAGI 服务：

```bash
curl -O https://raw.githubusercontent.com/fulaoaz/pentagi/zh-CN/docker-compose.yml
docker compose up -d
```

访问 [localhost:8443](https://localhost:8443) 即可打开 PentAGI 网页界面（默认账号为 `admin@pentagi.com` / `admin`）

#### 网页界面账号

PentAGI 的登录页不开放自助注册。全新安装会创建默认的本地管理员账号：

- **邮箱**：`admin@pentagi.com`
- **密码**：`admin`

首次登录后，请在正式使用前修改默认密码。若之后丢失管理员密码，可通过安装程序的维护菜单重置默认 `admin@pentagi.com` 账号的密码。

在多用户场景中，管理员登录后可以通过用户管理 REST API（`/api/v1/users/`）管理本地用户。实例启动后，可通过 `https://localhost:8443/api/v1/swagger/index.html` 打开 OpenAPI 界面。

> [!NOTE]
> 如果出现与 `pentagi-network`、`observability-network` 或 `langfuse-network` 有关的错误，请先运行 `docker-compose.yml` 创建这些网络，再运行 `docker-compose-langfuse.yml`、`docker-compose-graphiti.yml` 和 `docker-compose-observability.yml`，启用 Langfuse、Graphiti 和可观测性服务。
>
> 使用 PentAGI 至少需要配置一个语言模型提供商（OpenAI、Anthropic、Gemini、AWS Bedrock 或 Ollama）。AWS Bedrock 提供企业级的多厂商基础模型接入；如果你有足够的算力，Ollama 则可实现零成本的本地推理。搜索引擎的额外 API 密钥并非必需，但配置后通常能获得更好的结果。
>
> **使用先进模型进行全本地部署**：请参阅完整指南 [使用 vLLM 与 Qwen3.5-27B-FP8 运行 PentAGI](examples/guides/vllm-qwen35-27b-fp8.md)，了解生产级本地 LLM 部署方案。该配置在 4 张 RTX 5090 显卡上，提示处理速度约为 13,000 TPS，生成速度约为 650 TPS，可支持 12 个以上并发任务流，并且完全不依赖云端提供商。
>
> `LLM_SERVER_*` 系列环境变量属于实验性功能，未来可能变更。目前可用它们指定自定义 LLM 服务地址，并为所有智能体类型统一指定一个模型。
>
> `PROXY_URL` 是所有 LLM 提供商和外部搜索系统共用的全局代理地址，可统一控制这些服务对外网的访问。
>
> `docker-compose.yml` 以 root 用户运行 PentAGI 服务，因为容器管理需要访问 docker.sock。如果你使用 TCP/IP 方式连接 Docker 而非套接字文件，可以去掉 root 权限、改用默认的 `pentagi` 用户，以获得更好的安全性。

### 从外部网络访问 PentAGI

出于安全考虑，PentAGI 默认绑定到 `127.0.0.1`（仅本机可访问）。若要从网络中的其他机器访问 PentAGI，需要配置外部访问。

#### 配置步骤

1. **更新 `.env` 文件**，填入服务器的 IP 地址：

```bash
# 网络监听地址：允许外部设备连接
PENTAGI_LISTEN_IP=0.0.0.0
PENTAGI_LISTEN_PORT=8443

# 公网访问地址：请填写实际服务器 IP 或主机名
# 将 192.168.1.100 替换为服务器的实际 IP 地址
PUBLIC_URL=https://192.168.1.100:8443

# CORS 来源：列出所有需要访问 PentAGI 的 URL
# 同时保留本机访问所需的 localhost 和外部访问所需的服务器 IP
CORS_ORIGINS=https://localhost:8443,https://192.168.1.100:8443
```

> [!IMPORTANT]
> - 请把 `192.168.1.100` 替换为你服务器的实际 IP 地址
> - 不要在 `PUBLIC_URL` 或 `CORS_ORIGINS` 中使用 `0.0.0.0`，要填写实际 IP 地址
> - 建议在 `CORS_ORIGINS` 中同时包含 localhost 和服务器 IP，以便灵活访问

2. **重新创建容器**以应用变更：

```bash
docker compose down
docker compose up -d --force-recreate
```

3. **验证端口绑定：**

```bash
docker ps | grep pentagi
```

你应当看到 `0.0.0.0:8443->8443/tcp` 或 `:::8443->8443/tcp`。

如果看到的是 `127.0.0.1:8443->8443/tcp`，说明环境变量未生效。这种情况下请直接修改 `docker-compose.yml` 的第 31 行：

```yaml
ports:
  - "0.0.0.0:8443:8443"
```

然后重新创建容器。

4. **配置防火墙**，放通 8443 端口的入站连接：

```bash
# Ubuntu/Debian（使用 UFW）
sudo ufw allow 8443/tcp
sudo ufw reload

# CentOS/RHEL（使用 firewalld）
sudo firewall-cmd --permanent --add-port=8443/tcp
sudo firewall-cmd --reload
```

5. **访问 PentAGI：**

- **本机访问：** `https://localhost:8443`
- **网络访问：** `https://your-server-ip:8443`

> [!NOTE]
> 通过 IP 地址访问时，需要在浏览器中确认自签名 SSL 证书警告后继续访问。

---

### 使用 Podman 运行 PentAGI

PentAGI 完全支持 Podman，可将其作为 Docker 的替代方案。不过在 **Podman 无根（rootless）模式** 下，抓取服务需要额外配置，因为无根容器不能绑定小于 1024 的特权端口。

#### Podman 无根模式配置

抓取服务默认使用 443 端口（HTTPS），这是特权端口。在 Podman 无根模式下，需将抓取服务改到非特权端口：

**1. 编辑 `docker-compose.yml`** —— 修改 `scraper` 服务（大约在第 199 行）：

```yaml
scraper:
  image: vxcontrol/scraper:latest
  restart: unless-stopped
  container_name: scraper
  hostname: scraper
  expose:
    - 3000/tcp  # Changed from 443 to 3000
  ports:
    - "${SCRAPER_LISTEN_IP:-127.0.0.1}:${SCRAPER_LISTEN_PORT:-9443}:3000"  # Map to port 3000
  environment:
    - MAX_CONCURRENT_SESSIONS=${LOCAL_SCRAPER_MAX_CONCURRENT_SESSIONS:-10}
    - USERNAME=${LOCAL_SCRAPER_USERNAME:-someuser}
    - PASSWORD=${LOCAL_SCRAPER_PASSWORD:-somepass}
  logging:
    options:
      max-size: 50m
      max-file: "7"
  volumes:
    - scraper-ssl:/usr/src/app/ssl
  networks:
    - pentagi-network
  shm_size: 2g
```

**2. 更新 `.env` 文件** —— 将抓取服务地址改为使用 HTTP 和 3000 端口：

```bash
# Podman 无根模式下的抓取服务配置
SCRAPER_PRIVATE_URL=http://someuser:somepass@scraper:3000/
LOCAL_SCRAPER_USERNAME=someuser
LOCAL_SCRAPER_PASSWORD=somepass
```

> [!IMPORTANT]
> 使用 Podman 时的关键改动：
> - `SCRAPER_PRIVATE_URL` 使用 **HTTP** 而非 HTTPS
> - 使用 **3000** 端口而非 443
> - 将容器内部的 `expose` 改为 `3000/tcp`
> - 将端口映射的目标端口从 `443` 改为 `3000`

**3. 重新创建容器：**

```bash
podman-compose down
podman-compose up -d --force-recreate
```

**4. 测试抓取服务连通性：**

```bash
# 在 pentagi 容器内测试
podman exec -it pentagi wget -O- "http://someuser:somepass@scraper:3000/html?url=http://example.com"
```

如果输出了 HTML 内容，说明抓取服务工作正常。

#### Podman 有根模式

如果以有根模式（使用 sudo）运行 Podman，可直接使用默认配置，无需改动。抓取服务会按预期在 443 端口工作。

#### 与 Docker 的兼容性

上述 Podman 配置与 Docker 完全兼容。使用非特权端口的方案在两种容器运行时上的行为一致。

### 助手配置

PentAGI 允许你为助手配置默认行为：

| 变量                   | 默认值  | 说明                                 |
| ---------------------- | ------- | ------------------------------------ |
| `ASSISTANT_USE_AGENTS` | `false` | 新建助手时是否默认启用智能体委派     |

`ASSISTANT_USE_AGENTS` 决定在界面中新建助手时，“Use Agents（使用智能体）”开关的初始状态：
- `false`（默认）：新建助手时默认关闭智能体委派
- `true`：新建助手时默认启用智能体委派

创建或编辑助手时，用户随时可以在界面中切换“Use Agents（使用智能体）”开关。该环境变量只控制初始默认值。

## 登录后如何使用 PentAGI

服务启动并登录网页界面后，最快的上手方式是从 Flows（任务流）开始。

### 1. 创建第一个任务流

1. 在侧边栏中打开 **Flows（任务流）**。
2. 点击 **New Flow（新建任务流）**。
3. 根据目标选择运行模式：
   - **Automation（自动化）**：PentAGI 围绕测试目标自主完成整个流程
   - **Assistant（助手）**：适合通过交互逐步引导调查。在此模式下，还可以开启 **Use Agents（使用智能体）**，让 PentAGI 将复杂调查中的子任务委派给专职子智能体。
4. 选择此任务流使用的 LLM 提供商。
5. 在消息框中用自然语言描述目标和测试目的。

第一条提示词最好包含：

- 目标系统或 URL
- 所需的评估类型
- 测试范围限制或测试规则
- 预期结果，例如漏洞报告或对某项假设的验证

示例：

```text
评估 https://target.example 中常见的 Web 应用漏洞，重点检查身份认证、文件处理和注入问题。测试范围仅限给定目标，并汇总已确认的问题及复现步骤。
```

仅测试你拥有或明确获准评估的系统。可接受使用要求参见 [EULA.md](EULA.md)。

### 2. 使用模板复用任务流

新建任务流表单提供模板选择器，可以用已保存的任务流模板预填消息框。需要反复执行类似评估时，可直接复用这些模板。

- 如果已在 **Templates（模板）** 中保存模板，可直接选用
- 如需可直接使用的 Web 测试基线，可从 [`examples/prompts/base_web_pentest.md`](examples/prompts/base_web_pentest.md) 中的示例提示词开始
- 启动任务流前，按实际情况调整目标、范围和约束条件

模板只是起点。PentAGI 不要求使用特殊语法，只要目标和目的清楚，直接使用自然语言说明即可。

### 3. 监控执行过程并查看结果

提交任务流后，PentAGI 会自动打开任务流页面。

- 在任务流主视图中查看消息、智能体活动和任务进度
- 在运行过程中检查工具活动和终端输出
- 查看生成的任务和子任务，了解 PentAGI 正在执行的操作

任务流产生足够的结果后，可以通过页面上的 **Report（报告）** 菜单：

- 在网页中打开报告
- 将生成的报告复制到剪贴板
- 下载 Markdown 格式的报告
- 下载 PDF 格式的报告

### 4. 通过助手视图调整进行中的任务流

每个任务流都有 **Assistant（助手）** 视图，可通过交互进行引导。当自动化执行遇到需要人工判断的情况时，可以直接介入，无需从头重启任务流。

- 如需在改动前查看当前状态，请打开同一任务流的 **Assistant（助手）** 视图。
- 可以通过助手检查任务流状态、停止当前任务、提交后续指令，或在下一步执行前调整尚未运行的计划子任务。
- 该视图是当前任务流的明确控制入口，并非隐藏的后台队列。需要改变方向时，请清楚说明新要求，并确保新指令仍在当前测试范围内。
- 它适合用来澄清范围、根据中间结果调整优先级，或响应自动化检查点，同时保留任务流的其余上下文。

### 5. 管理任务流专属文件

每个任务流页面都有独立的 **Files（文件）** 标签页。文件仅属于其父任务流，存放在主机的 `{dataDir}/flow-{id}-data/` 目录中，不会进入其他任务流。

该标签页包含三类文件来源：

- **Uploads（上传文件）**（`uploads/`）：从网页界面上传的文件。可以使用 **Upload files（上传文件）** 操作，也可以直接将文件拖放到 Files 标签页。智能体容器运行期间，上传的文件还会同步到容器内的 `/work/uploads/`，供智能体使用常规 shell 工具读取。
- **Resources（资源）**（`resources/`）：通过 **Attach resources from library（从资源库附加资源）** 从用户资源库附加的文件。这些资源会复制到任务流中，并同步到运行中容器的 `/work/resources/`。
- **Container（容器）**（`container/`）：通过 **Pull file or directory from container（从容器拉取文件或目录）** 从运行中的智能体容器获取的快照。这些文件在任务流端为只读，且不会传回容器。

Files 标签页为每个文件提供 **Download（下载）**、**Copy path（复制路径）**、**Save as resource（另存为资源）** 和 **Delete（删除）** 操作。“另存为资源”会把任务流文件加入可复用的用户资源库。容器未运行时，Pull（拉取）操作会被禁用，并显示“Container is not running（容器未运行）”提示。

`{{.UserFiles}}` 模板变量会将上传文件和附加资源自动列入智能体的系统提示词，并生成紧凑的 `<task_files>` XML 块，其中包含 `<uploads>` 和 `<resources>` 子节。这样，助手和自动化智能体可以直接按路径引用文件，无需把文件内容粘贴到对话中。容器快照只在界面中可见，不会自动加入提示词。

目前有以下限制：

- 单个上传文件最大为 300 MB；每次上传最多包含 1000 个文件，总大小不超过 2 GB。文件名最长 255 字节，约等于 255 个 ASCII 字符；非 ASCII 字符通常占用多个字节。
- 上传文件和资源会分别同步到运行中容器的固定路径 `/work/uploads/` 和 `/work/resources/`。写入容器其他路径的文件不会自动同步回任务流文件模型。可以从容器内任意指定路径拉取快照，例如 `/etc/...`；快照会缓存在任务流端的 `container/` 目录下，但不会再推送回容器。
- 容器快照只记录拉取时的状态。在界面中编辑快照不会写回运行中的容器。
- 目前删除任务流只会移除任务流记录及其长期记忆条目，不会归档或删除磁盘上的 `flow-{id}-data/` 目录。如需回收空间，运维人员仍需手动清理该数据目录。

初次测试时，建议从范围较小的目标和单一、明确的目的开始。这样更容易检查输出，也便于在执行更大规模的评估前完善提示词。

## API 访问

PentAGI 同时提供 REST 与 GraphQL API，可通过程序调用其功能，并将渗透测试任务流集成到自动化管道、CI/CD 流程和自定义应用中。

### 生成 API 令牌

API 令牌通过 PentAGI 网页界面管理：

1. 在网页界面中打开 **Settings -> API Tokens（设置 -> API 令牌）**
2. 点击 **Create Token（创建令牌）** 生成新的 API 令牌
3. 设置令牌属性：
   - **Name（名称）**（可选）：便于识别令牌的名称
   - **Expiration Date（到期日期）**：令牌的到期时间，最短 1 分钟，最长 3 年
4. 点击 **Create（创建）** 后**立即复制令牌**。出于安全考虑，令牌只会显示一次
5. 在 API 请求中将该令牌用作 Bearer 令牌

每个令牌都与用户账号关联，并继承该账号所属角色的权限。

### 使用 API 令牌

在 HTTP 请求的 `Authorization` 请求头中加入 API 令牌：

```bash
# GraphQL API 示例
curl -X POST https://your-pentagi-instance:8443/api/v1/graphql \
  -H "Authorization: Bearer YOUR_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"query": "{ flows { id title status } }"}'

# REST API 示例
curl https://your-pentagi-instance:8443/api/v1/flows \
  -H "Authorization: Bearer YOUR_API_TOKEN"
```

### 浏览和测试 API

PentAGI 提供交互式文档，可用于浏览和测试 API 端点：

#### GraphQL Playground

通过 `https://your-pentagi-instance:8443/api/v1/graphql/playground` 打开 GraphQL Playground。

1. 点击底部的 **HTTP Headers（HTTP 请求头）** 标签
2. 添加认证请求头：
   ```json
   {
     "Authorization": "Bearer YOUR_API_TOKEN"
   }
   ```
3. 以交互方式浏览 schema、运行查询并测试 mutation

#### Swagger UI

通过 `https://your-pentagi-instance:8443/api/v1/swagger/index.html` 打开 REST API 文档。

1. 点击 **Authorize（认证）** 按钮
2. 按以下格式输入令牌：`Bearer YOUR_API_TOKEN`
3. 点击 **Authorize（认证）** 应用令牌
4. 直接在 Swagger UI 中测试端点

### 生成 API 客户端

可以使用 PentAGI 附带的 schema 文件，为所需编程语言生成类型安全的 API 客户端：

#### GraphQL 客户端

可通过以下方式获取 GraphQL schema：
- **网页界面**：进入 Settings（设置）下载 `schema.graphqls`
- **仓库文件**：`backend/pkg/graph/schema.graphqls`

可使用以下工具生成客户端：
- **GraphQL Code Generator**（JavaScript/TypeScript）：[https://the-guild.dev/graphql/codegen](https://the-guild.dev/graphql/codegen)
- **genqlient**（Go）：[https://github.com/Khan/genqlient](https://github.com/Khan/genqlient)
- **Apollo iOS**（Swift）：[https://www.apollographql.com/docs/ios](https://www.apollographql.com/docs/ios)

#### REST API 客户端

可通过以下方式获取 OpenAPI 规范：
- **Swagger JSON**：`https://your-pentagi-instance:8443/api/v1/swagger/doc.json`
- **Swagger YAML**：仓库中的 `backend/pkg/server/docs/swagger.yaml`

可使用以下工具生成客户端：
- **OpenAPI Generator**：[https://openapi-generator.tech](https://openapi-generator.tech)
  ```bash
  openapi-generator-cli generate \
    -i https://your-pentagi-instance:8443/api/v1/swagger/doc.json \
    -g python \
    -o ./pentagi-client
  ```

- **Swagger Codegen**：[https://github.com/swagger-api/swagger-codegen](https://github.com/swagger-api/swagger-codegen)
  ```bash
  swagger-codegen generate \
    -i https://your-pentagi-instance:8443/api/v1/swagger/doc.json \
    -l typescript-axios \
    -o ./pentagi-client
  ```

- **swagger-typescript-api**（TypeScript）：[https://github.com/acacode/swagger-typescript-api](https://github.com/acacode/swagger-typescript-api)
  ```bash
  npx swagger-typescript-api \
    -p https://your-pentagi-instance:8443/api/v1/swagger/doc.json \
    -o ./src/api \
    -n pentagi-api.ts
  ```

### API 使用示例

<details>
<summary><b>创建新任务流（GraphQL）</b></summary>

```graphql
mutation CreateFlow {
  createFlow(
    modelProvider: "openai"
    input: "测试 https://example.com 的安全性"
  ) {
    id
    title
    status
    createdAt
  }
}
```

</details>

<details>
<summary><b>列出任务流（REST API）</b></summary>

```bash
curl https://your-pentagi-instance:8443/api/v1/flows \
  -H "Authorization: Bearer YOUR_API_TOKEN" \
  | jq '.flows[] | {id, title, status}'
```

</details>

<details>
<summary><b>Python 客户端示例</b></summary>

```python
import requests

class PentAGIClient:
    def __init__(self, base_url, api_token):
        self.base_url = base_url
        self.headers = {
            "Authorization": f"Bearer {api_token}",
            "Content-Type": "application/json"
        }
    
    def create_flow(self, provider, target):
        query = """
        mutation CreateFlow($provider: String!, $input: String!) {
          createFlow(modelProvider: $provider, input: $input) {
            id
            title
            status
          }
        }
        """
        response = requests.post(
            f"{self.base_url}/api/v1/graphql",
            json={
                "query": query,
                "variables": {
                    "provider": provider,
                    "input": target
                }
            },
            headers=self.headers
        )
        return response.json()
    
    def get_flows(self):
        response = requests.get(
            f"{self.base_url}/api/v1/flows",
            headers=self.headers
        )
        return response.json()

# 使用方法
client = PentAGIClient(
    "https://your-pentagi-instance:8443",
    "your_api_token_here"
)

# 创建新任务流
flow = client.create_flow("openai", "扫描 https://example.com 中的漏洞")
print(f"已创建任务流：{flow}")

# 列出所有任务流
flows = client.get_flows()
print(f"任务流总数：{len(flows['flows'])}")
```

</details>

<details>
<summary><b>TypeScript 客户端示例</b></summary>

```typescript
import axios, { AxiosInstance } from 'axios';

interface Flow {
  id: string;
  title: string;
  status: string;
  createdAt: string;
}

class PentAGIClient {
  private client: AxiosInstance;

  constructor(baseURL: string, apiToken: string) {
    this.client = axios.create({
      baseURL: `${baseURL}/api/v1`,
      headers: {
        'Authorization': `Bearer ${apiToken}`,
        'Content-Type': 'application/json',
      },
    });
  }

  async createFlow(provider: string, input: string): Promise<Flow> {
    const query = `
      mutation CreateFlow($provider: String!, $input: String!) {
        createFlow(modelProvider: $provider, input: $input) {
          id
          title
          status
          createdAt
        }
      }
    `;

    const response = await this.client.post('/graphql', {
      query,
      variables: { provider, input },
    });

    return response.data.data.createFlow;
  }

  async getFlows(): Promise<Flow[]> {
    const response = await this.client.get('/flows');
    return response.data.flows;
  }

  async getFlow(flowId: string): Promise<Flow> {
    const response = await this.client.get(`/flows/${flowId}`);
    return response.data;
  }
}

// 使用方法
const client = new PentAGIClient(
  'https://your-pentagi-instance:8443',
  'your_api_token_here'
);

// 创建新任务流
const flow = await client.createFlow(
  'openai',
  '对 https://example.com 执行渗透测试'
);
console.log('已创建任务流：', flow);

// 列出所有任务流
const flows = await client.getFlows();
console.log(`任务流总数：${flows.length}`);
```

</details>

### 安全最佳实践

使用 API 令牌时：

- **切勿将令牌提交到版本控制系统**：请使用环境变量或密钥管理系统
- **定期轮换令牌**：设置合适的到期日期，并定期创建新令牌
- **为不同应用分配独立令牌**：需要收回访问权限时更容易单独吊销
- **监控令牌使用情况**：在 Settings（设置）页面查看 API 令牌活动
- **吊销不再使用的令牌**：禁用或删除不再需要的令牌
- **仅使用 HTTPS**：不要通过未加密连接发送 API 令牌

### 令牌管理

- **查看令牌**：在 Settings -> API Tokens（设置 -> API 令牌）中查看所有有效令牌
- **编辑令牌**：更新令牌名称或吊销令牌
- **删除令牌**：永久删除令牌，此操作不可撤销
- **令牌 ID**：每个令牌都有唯一 ID，可复制留作参考

令牌列表显示以下信息：
- 令牌名称（如已填写）
- 令牌 ID（唯一标识符）
- 状态（有效/已吊销/已过期）
- 创建日期
- 到期日期

### 自定义 LLM 提供商配置

通过 `LLM_SERVER_*` 变量使用自定义 LLM 提供商时，可以调整请求中的推理参数格式。

> [!TIP]
> 如果要搭建可用于生产环境的本地部署，建议通过 **vLLM** 运行 **Qwen3.5-27B-FP8**，以获得较好的性能。请参阅[完整部署指南](examples/guides/vllm-qwen35-27b-fp8.md)，其中包含硬件要求、[思考模式](examples/configs/vllm-qwen3.5-27b-fp8.provider.yml)与[非思考模式](examples/configs/vllm-qwen3.5-27b-fp8-no-think.provider.yml)的配置模板，以及在 4 张 RTX 5090 GPU 上提示词处理吞吐量达到 1.3 万 TPS 的性能测试结果。

| 变量                            | 默认值  | 说明                                                                                    |
| ------------------------------- | ------- | --------------------------------------------------------------------------------------- |
| `LLM_SERVER_URL`                |         | 自定义 LLM API 端点的基础 URL                                                           |
| `LLM_SERVER_KEY`                |         | 自定义 LLM 提供商的 API 密钥                                                            |
| `LLM_SERVER_MODEL`              |         | 默认模型（可在提供商配置中覆盖）                                                        |
| `LLM_SERVER_CONFIG_PATH`        |         | 为各智能体指定模型的 YAML 配置文件路径                                                  |
| `LLM_SERVER_PROVIDER`           |         | 模型名称中的提供商前缀（例如 LiteLLM 代理使用的 `openrouter`、`deepseek`）              |
| `LLM_SERVER_LEGACY_REASONING`   | `false` | 控制 API 请求中的推理参数格式                                                           |
| `LLM_SERVER_PRESERVE_REASONING` | `false` | 在多轮对话中保留推理内容（部分提供商要求开启）                                          |

使用 **LiteLLM 代理**时，`LLM_SERVER_PROVIDER` 很有用，因为 LiteLLM 会在模型名称前添加提供商前缀。例如，通过 LiteLLM 连接 Moonshot API 时，`kimi-2.5` 之类的模型名称会变成 `moonshot/kimi-2.5`。设置 `LLM_SERVER_PROVIDER=moonshot` 后，同一份提供商配置文件既可直连 API，也可通过 LiteLLM 代理使用，无需修改。

`LLM_SERVER_LEGACY_REASONING` 决定如何向 LLM 发送推理参数：
- `false`（默认）：使用新版格式，以带有 `max_tokens` 参数的结构化对象发送推理设置
- `true`：使用旧版格式，通过字符串类型的 `reasoning_effort` 参数发送

不同 LLM 提供商对 API 请求中的推理格式要求可能不同。如果自定义提供商返回与推理参数有关的错误，请尝试切换此设置。

`LLM_SERVER_PRESERVE_REASONING` 决定是否在多轮对话中保留推理内容：
- `false`（默认）：不在对话历史中保留推理内容
- `true`：保留推理内容，并在后续 API 调用中一并发送

Moonshot 等部分 LLM 提供商要求多轮对话包含推理内容，否则会返回 `thinking is enabled but reasoning_content is missing in assistant tool call message` 一类错误。如果所用提供商有此要求，请开启该设置。

### Ollama 提供商配置

PentAGI 支持通过 Ollama 在本地运行 LLM 推理（免 API 调用费，数据更私密），也支持带免费套餐的托管服务 Ollama Cloud。

#### 配置变量

| 变量                                | 默认值      | 说明                                      |
| ----------------------------------- | ----------- | ----------------------------------------- |
| `OLLAMA_SERVER_URL`                 |             | Ollama 服务器或 Ollama Cloud 的 URL        |
| `OLLAMA_SERVER_API_KEY`             |             | 用于 Ollama Cloud 认证的 API 密钥          |
| `OLLAMA_SERVER_MODEL`               |             | 默认推理模型                              |
| `OLLAMA_SERVER_CONFIG_PATH`         |             | 自定义智能体配置文件的路径                |
| `OLLAMA_SERVER_PULL_MODELS_TIMEOUT` | `600`       | 模型下载超时时间（秒）                    |
| `OLLAMA_SERVER_PULL_MODELS_ENABLED` | `false`     | 启动时自动下载模型                        |
| `OLLAMA_SERVER_LOAD_MODELS_ENABLED` | `false`     | 向服务器查询可用模型                      |

#### Ollama Cloud 配置

Ollama Cloud 提供托管推理服务，可选择免费套餐或按需扩展的付费套餐。

**免费套餐设置（单模型）**

```bash
# 免费套餐一次只能使用一个模型
OLLAMA_SERVER_URL=https://ollama.com
OLLAMA_SERVER_API_KEY=your_ollama_cloud_api_key
OLLAMA_SERVER_MODEL=gpt-oss:120b  # 示例：OpenAI OSS 120B 模型
```

**付费套餐设置（多模型及预置配置）**

付费套餐支持同时使用多个模型，可直接采用预置的 Ollama Cloud 配置：

```bash
# 使用 Docker 镜像内置的 Ollama Cloud 配置
OLLAMA_SERVER_URL=https://ollama.com
OLLAMA_SERVER_API_KEY=your_ollama_cloud_api_key
OLLAMA_SERVER_CONFIG_PATH=/opt/pentagi/conf/ollama-cloud.provider.yml
```

预置的 `ollama-cloud.provider.yml` 已针对各类智能体分配模型：
- **Simple/Assistant（简单任务/助手）**：`nemotron-3-super:cloud`，快速通用模型
- **Primary Agent（主智能体）**：`qwen3-coder-next:cloud`，支持高强度推理
- **Coder/Pentester（编码/渗透测试）**：`qwen3-coder-next:cloud`，面向编码任务的专用模型
- **Searcher（搜索）**：`qwen3.5:397b-cloud`，上下文窗口较大，适合收集信息
- **Refiner/Refactor（优化/重构）**：`glm-5:cloud`，用于高质量文本优化
- **Adviser/Enricher（顾问/内容补充）**：`minimax-m2.7:cloud`，适合高效处理建议类任务
- **Installer（安装）**：`devstral-2:123b-cloud`，用于安装和设置任务

**自定义配置（高级）**

如需自行配置智能体，请从宿主机文件系统挂载自定义文件：

```bash
# 使用自定义提供商配置
OLLAMA_SERVER_URL=https://ollama.com
OLLAMA_SERVER_API_KEY=your_ollama_cloud_api_key
OLLAMA_SERVER_CONFIG_PATH=/opt/pentagi/conf/ollama.provider.yml

# 从宿主机文件系统挂载自定义配置（在 .env 或 docker-compose 覆盖文件中设置）
PENTAGI_OLLAMA_SERVER_CONFIG_PATH=/path/on/host/my-ollama-config.yml
```

环境变量 `PENTAGI_OLLAMA_SERVER_CONFIG_PATH` 会将宿主机上的配置文件映射到容器内的 `/opt/pentagi/conf/ollama.provider.yml`。

**自定义配置示例**（`my-ollama-config.yml`）：

```yaml
primary_agent:
  model: "qwen3-coder-next:cloud"
  temperature: 1.0
  top_p: 0.9
  max_tokens: 32768
  reasoning:
    effort: high

coder:
  model: "qwen3-coder:32b"
  temperature: 1.0
  max_tokens: 20480
```

#### 本地 Ollama 配置

自托管 Ollama 实例的配置如下：

```bash
# 本地 Ollama 基本设置
OLLAMA_SERVER_URL=http://localhost:11434
OLLAMA_SERVER_MODEL=llama3.1:8b-instruct-q8_0

# 生产环境设置，启用自动拉取和模型发现
OLLAMA_SERVER_URL=http://ollama-server:11434
OLLAMA_SERVER_PULL_MODELS_ENABLED=true
OLLAMA_SERVER_PULL_MODELS_TIMEOUT=900
OLLAMA_SERVER_LOAD_MODELS_ENABLED=true

# 使用 Docker 镜像中的预置配置
OLLAMA_SERVER_CONFIG_PATH=/opt/pentagi/conf/ollama-llama318b.provider.yml
# 或
OLLAMA_SERVER_CONFIG_PATH=/opt/pentagi/conf/ollama-qwen332b-fp16-tc.provider.yml
# 或
OLLAMA_SERVER_CONFIG_PATH=/opt/pentagi/conf/ollama-qwq32b-fp16-tc.provider.yml
```

**性能注意事项：**

- **模型发现**（`OLLAMA_SERVER_LOAD_MODELS_ENABLED=true`）：查询 Ollama API 会使启动时间增加 1～2 秒
- **自动拉取**（`OLLAMA_SERVER_PULL_MODELS_ENABLED=true`）：首次启动需要下载模型，可能耗时数分钟
- **拉取超时**（`OLLAMA_SERVER_PULL_MODELS_TIMEOUT=900`）：以秒为单位设置 15 分钟超时
- **静态配置**：关闭上述两个开关，在配置文件中指定模型，可以缩短启动时间

#### 创建扩展上下文的自定义 Ollama 模型

PentAGI 所需的上下文窗口大于 Ollama 默认配置。需要通过 Modelfile 增大 `num_ctx` 参数并创建自定义模型。常规智能体任务约消耗 64K 令牌，PentAGI 将上下文大小设为 110K，以便为复杂渗透测试场景预留余量。

**注意**：`num_ctx` 只能在通过 Modelfile 创建模型时设置；模型创建后无法修改，也无法在运行时覆盖。

##### 示例：扩展上下文的 Qwen3 32B FP16

创建名为 `Modelfile_qwen3_32b_fp16_tc` 的 Modelfile：

```dockerfile
FROM qwen3:32b-fp16
PARAMETER num_ctx 110000
PARAMETER temperature 0.3
PARAMETER top_p 0.8
PARAMETER min_p 0.0
PARAMETER top_k 20
PARAMETER repeat_penalty 1.1
```

构建自定义模型：

```bash
ollama create qwen3:32b-fp16-tc -f Modelfile_qwen3_32b_fp16_tc
```

##### 示例：扩展上下文的 QwQ 32B FP16

创建名为 `Modelfile_qwq_32b_fp16_tc` 的 Modelfile：

```dockerfile
FROM qwq:32b-fp16
PARAMETER num_ctx 110000
PARAMETER temperature 0.2
PARAMETER top_p 0.7
PARAMETER min_p 0.0
PARAMETER top_k 40
PARAMETER repeat_penalty 1.2
```

构建自定义模型：

```bash
ollama create qwq:32b-fp16-tc -f Modelfile_qwq_32b_fp16_tc
```

> **注意**：QwQ 32B FP16 推理大约需要 **71.3 GB 显存**。使用前请确认系统有足够的 GPU 显存。

Docker 镜像的 `/opt/pentagi/conf/` 目录中包含预置提供商配置文件 `ollama-qwen332b-fp16-tc.provider.yml` 和 `ollama-qwq32b-fp16-tc.provider.yml`，其中已引用上述自定义模型。

### OpenAI 提供商配置

PentAGI 可接入 OpenAI 的多个模型，包括支持长思维链的高级推理模型、工具集成能力更强的智能体模型，以及面向安全工程的代码专用模型。

#### 配置变量

| 变量                 | 默认值                      | 说明                        |
| -------------------- | --------------------------- | --------------------------- |
| `OPEN_AI_KEY`        |                             | OpenAI 服务的 API 密钥      |
| `OPEN_AI_SERVER_URL` | `https://api.openai.com/v1` | OpenAI API 端点             |

#### 配置示例

```bash
# OpenAI 基本设置
OPEN_AI_KEY=your_openai_api_key
OPEN_AI_SERVER_URL=https://api.openai.com/v1

# 通过代理连接以提高安全性
OPEN_AI_KEY=your_openai_api_key
PROXY_URL=http://your-proxy:8080
```

#### 支持的模型

PentAGI 支持 31 个 OpenAI 模型，可使用工具调用、流式输出、推理模式和提示词缓存。标有 `*` 的模型用于默认配置。

**GPT-5.2 系列：最新旗舰智能体模型（2025 年 12 月）**

| 模型 ID               | 推理     | 价格（输入/输出/缓存）     | 适用场景                                        |
| --------------------- | -------- | -------------------------- | ----------------------------------------------- |
| `gpt-5.2`*            | ✅        | $1.75/$14.00/$0.18         | 最新旗舰模型，增强了推理和工具集成能力，适合自主安全研究 |
| `gpt-5.2-pro`         | ✅        | $21.00/$168.00/$0.00       | 高端版本，智能体编码能力更强，适合任务关键型安全研究和零日漏洞发现 |
| `gpt-5.2-codex`       | ✅        | $1.75/$14.00/$0.18         | 高级代码专用模型，支持上下文压缩，网络安全能力较强 |

**GPT-5/5.1 系列：高级智能体模型**

| 模型 ID               | 推理     | 价格（输入/输出/缓存）     | 适用场景                                        |
| --------------------- | -------- | -------------------------- | ----------------------------------------------- |
| `gpt-5`               | ✅        | $1.25/$10.00/$0.13         | 旗舰智能体模型，具有高级推理能力，适合自主安全研究和利用链开发 |
| `gpt-5.1`             | ✅        | $1.25/$10.00/$0.13         | 增强型智能体模型，支持自适应推理和较强的工具协同，适合综合渗透测试 |
| `gpt-5-pro`           | ✅        | $15.00/$120.00/$0.00       | 高端版本，大幅改进推理并减少幻觉，适合关键安全操作 |
| `gpt-5-mini`          | ✅        | $0.25/$2.00/$0.03          | 兼顾速度与智能，适合自动化漏洞分析和漏洞利用生成 |
| `gpt-5-nano`          | ✅        | $0.05/$0.40/$0.01          | 速度最快，适合高吞吐扫描、侦察和批量漏洞检测 |

**GPT-5/5.1 Codex 系列：代码专用模型**

| 模型 ID               | 推理     | 价格（输入/输出/缓存）     | 适用场景                                        |
| --------------------- | -------- | -------------------------- | ----------------------------------------------- |
| `gpt-5.1-codex-max`   | ✅        | $1.25/$10.00/$0.13         | 增强推理能力，适合复杂编码、CVE 发现和系统化漏洞利用开发 |
| `gpt-5.1-codex`       | ✅        | $1.25/$10.00/$0.13         | 标准代码优化模型，推理能力较强，适合漏洞利用生成和漏洞分析 |
| `gpt-5-codex`         | ✅        | $1.25/$10.00/$0.13         | 基础代码专用模型，适合漏洞扫描和基本漏洞利用生成 |
| `gpt-5.1-codex-mini`  | ✅        | $0.25/$2.00/$0.03          | 小型高性能模型，容量为原来的 4 倍，适合快速漏洞检测 |
| `codex-mini-latest`   | ✅        | $1.50/$6.00/$0.38          | 最新小型代码模型，适合自动代码审查和基本漏洞分析 |

**GPT-4.1 系列：增强智能模型**

| 模型 ID               | 推理     | 价格（输入/输出/缓存）     | 适用场景                                        |
| --------------------- | -------- | -------------------------- | ----------------------------------------------- |
| `gpt-4.1`             | ❌        | $2.00/$8.00/$0.50          | 增强型旗舰模型，函数调用能力出色，适合复杂威胁分析和高级漏洞利用开发 |
| `gpt-4.1-mini`*       | ❌        | $0.40/$1.60/$0.10          | 兼顾性能与效率，适合常规安全评估和自动代码分析 |
| `gpt-4.1-nano`        | ❌        | $0.10/$0.40/$0.03          | 超高速轻量模型，适合批量安全扫描、快速侦察和持续监控 |

**GPT-4o 系列：多模态旗舰模型**

| 模型 ID               | 推理     | 价格（输入/输出/缓存）     | 适用场景                                        |
| --------------------- | -------- | -------------------------- | ----------------------------------------------- |
| `gpt-4o`              | ❌        | $2.50/$10.00/$1.25         | 支持视觉的多模态旗舰模型，适合图像分析、网页界面评估和多工具编排 |
| `gpt-4o-mini`         | ❌        | $0.15/$0.60/$0.08          | 小型多模态模型，函数调用能力较强，适合高频扫描和低成本批量操作 |

**o 系列：高级推理模型**

| 模型 ID               | 推理     | 价格（输入/输出/缓存）     | 适用场景                                        |
| --------------------- | -------- | -------------------------- | ----------------------------------------------- |
| `o4-mini`*            | ✅        | $1.10/$4.40/$0.28          | 新一代快速推理模型，适合有条理的安全评估和系统化漏洞利用开发 |
| `o3`*                 | ✅        | $2.00/$8.00/$0.50          | 高级推理模型，适合多阶段攻击链和深度漏洞分析 |
| `o3-mini`             | ✅        | $1.10/$4.40/$0.55          | 小型推理模型，支持扩展思考，适合逐步制定攻击方案和漏洞逻辑串联 |
| `o1`                  | ✅        | $15.00/$60.00/$7.50        | 深度推理模型，适合高级渗透测试和新型漏洞利用研究 |
| `o3-pro`              | ✅        | $20.00/$80.00/$0.00        | 高级推理模型，比 o1-pro 便宜 80%，适合零日漏洞研究和关键安全调查 |
| `o1-pro`              | ✅        | $150.00/$600.00/$0.00      | 上一代高端推理模型，适合详尽安全分析和任务关键型难题 |

**价格**：按每 100 万令牌计费。推理模型的输出价格包含思考令牌。

> [!WARNING]
> **GPT-5* 模型：需要可信访问权限**
>
> 所有 GPT-5 系列模型（`gpt-5`、`gpt-5.1`、`gpt-5.2`、`gpt-5-pro`、`gpt-5.2-pro` 及全部 Codex 变体）在 PentAGI 中运行时都可能**不稳定**；如果未获得已验证的访问权限，还可能触发 OpenAI 的网络安全保护机制。
>
> **如需稳定使用 GPT-5* 模型：**
> 1. **个人用户**：在 [chatgpt.com/cyber](https://chatgpt.com/cyber) 完成身份验证
> 2. **企业团队**：通过 OpenAI 客户代表申请可信访问权限
> 3. **安全研究人员**：申请[网络安全资助计划](https://openai.com/form/cybersecurity-grant-program/)（提供总额 1,000 万美元的 API 额度）
>
> **无需验证的建议替代方案：**
> - 推理任务使用 `o-series` 模型（o3、o4-mini、o1）
> - 通用智能和函数调用使用 `gpt-4.1` 系列
> - 所有 o 系列和 gpt-4.x 模型无需特殊访问权限即可稳定运行

**推理强度级别**：
- **High（高）**：最大推理深度（`refiner` 使用 high 强度的 o3）
- **Medium（中）**：平衡推理深度与开销（`primary_agent`、`assistant`、`reflector` 使用 medium 强度的 o4-mini/o3）
- **Low（低）**：高效执行目标明确的推理任务（`coder`、`installer`、`pentester` 使用 low 强度的 o3/o4-mini；`adviser` 使用 low 强度的 gpt-5.2）

**主要特性**：
- **扩展推理**：o 系列模型通过思维链处理复杂安全分析
- **智能体能力**：GPT-5/5.1/5.2 系列增强了工具集成和自主执行能力
- **提示词缓存**：重复使用上下文时降低成本，缓存价格为输入价格的 10%～50%
- **代码专用模型**：Codex 模型专门用于漏洞发现和漏洞利用开发
- **多模态支持**：GPT-4o 系列可执行基于视觉信息的安全评估
- **工具调用**：所有模型均提供可靠的函数调用，可编排渗透测试工具
- **流式输出**：为交互式任务流实时返回响应
- **实际成果**：这些模型已用于发现 CVE，并应用于真实安全场景

### Anthropic 提供商配置

PentAGI 可接入 Anthropic 的 Claude 模型，支持高级扩展思考、提示词缓存和安全机制，并能理解复杂的安全上下文。

#### 配置变量

| 变量                   | 默认值                         | 说明                           |
| ---------------------- | ------------------------------ | ------------------------------ |
| `ANTHROPIC_API_KEY`    |                                | Anthropic 服务的 API 密钥      |
| `ANTHROPIC_SERVER_URL` | `https://api.anthropic.com/v1` | Anthropic API 端点             |

#### 配置示例

```bash
# Anthropic 基本设置
ANTHROPIC_API_KEY=your_anthropic_api_key
ANTHROPIC_SERVER_URL=https://api.anthropic.com/v1

# 在安全要求较高的环境中通过代理连接
ANTHROPIC_API_KEY=your_anthropic_api_key
PROXY_URL=http://your-proxy:8080
```

> [!NOTE]
> **通过 Google Vertex AI 使用 Claude 模型**
>
> PentAGI 目前没有在 `.env` 中提供 Anthropic Claude 专用的 Google Vertex AI 配置，也没有单独的 Vertex AI API 密钥字段。现有 Anthropic 变量（`ANTHROPIC_API_KEY`、`ANTHROPIC_SERVER_URL`）用于直连 Anthropic API。目前支持通过以下方式接入 Claude：
>
> - **直连 Anthropic API**：使用 `ANTHROPIC_API_KEY` 和 `ANTHROPIC_SERVER_URL`（见上文）
> - **AWS Bedrock**：使用 `BEDROCK_*` 变量（参见 [AWS Bedrock 提供商配置](#aws-bedrock-提供商配置)）
>
> 如果当前需要使用 Vertex AI，可以通过兼容 OpenAI 的代理或网关暴露 Vertex AI。代理或网关必须把 Vertex AI 调用转换为 Chat Completions 格式，同时保留 PentAGI 所依赖的聊天和工具调用行为。然后通过 `LLM_SERVER_URL`、`LLM_SERVER_KEY` 和 `LLM_SERVER_MODEL` 将自定义 LLM 提供商指向该网关。这种接入方式的可靠性取决于所选网关。

#### 支持的模型

PentAGI 支持 10 个 Claude 模型，可使用工具调用、流式输出、扩展思考、自适应思考和提示词缓存。标有 `*` 的模型用于默认配置。

**Claude 4 系列：最新模型（2025 至 2026 年）**

| 模型 ID                  | 思考     | 发布日期     | 价格（输入/输出/缓存读/写）    | 适用场景                                        |
| ------------------------ | -------- | ------------ | ------------------------------ | ----------------------------------------------- |
| `claude-opus-4-6`*       | ✅        | 2025 年 5 月 | $5.00/$25.00/$0.50/$6.25       | 智能水平最高，适合自主智能体和编码任务。扩展思考和自适应思考可用于复杂漏洞利用开发及多阶段攻击模拟 |
| `claude-sonnet-4-6`*     | ✅        | 2025 年 8 月 | $3.00/$15.00/$0.30/$3.75       | 速度与智能的平衡较好，并支持自适应思考。适合多阶段安全评估、智能漏洞分析和实时威胁搜寻 |
| `claude-haiku-4-5`*      | ✅        | 2025 年 10 月 | $1.00/$5.00/$0.10/$1.25       | 速度最快，智能水平接近前沿模型。适合高频扫描、实时监控和批量自动化测试 |

**旧版模型：仍受支持**

| 模型 ID                  | 思考     | 发布日期     | 价格（输入/输出/缓存读/写）    | 适用场景                                        |
| ------------------------ | -------- | ------------ | ------------------------------ | ----------------------------------------------- |
| `claude-sonnet-4-5`      | ✅        | 2025 年 9 月 | $3.00/$15.00/$0.30/$3.75       | 先进推理模型（已由 4-6 取代）。适合复杂渗透测试和高级威胁分析 |
| `claude-opus-4-5`        | ✅        | 2025 年 11 月 | $5.00/$25.00/$0.50/$6.25      | 高端推理模型（已由 opus-4-6 取代）。适合关键安全研究、零日漏洞发现和红队行动 |
| `claude-opus-4-1`        | ✅        | 2025 年 8 月 | $15.00/$75.00/$1.50/$18.75     | 高级推理模型（已被后续型号取代）。适合复杂渗透测试和高级威胁建模 |
| `claude-sonnet-4-0`      | ✅        | 2025 年 5 月 | $3.00/$15.00/$0.30/$3.75       | 高性能推理模型（已被后续型号取代）。适合复杂威胁建模和多工具协同 |
| `claude-opus-4-0`        | ✅        | 2025 年 5 月 | $15.00/$75.00/$1.50/$18.75     | 第一代 Opus（已被后续型号取代）。适合多步骤漏洞利用开发和自主渗透测试任务流 |

**弃用模型：请迁移到当前模型**

| 模型 ID                      | 思考     | 发布日期     | 价格（输入/输出/缓存读/写）    | 备注                                         |
| ---------------------------- | -------- | ------------ | ------------------------------ | -------------------------------------------- |
| `claude-3-haiku-20240307`    | ❌        | 2024 年 3 月 | $0.25/$1.25/$0.03/$0.30        | 将于 2026 年 4 月 19 日停用，请迁移到 claude-haiku-4-5 |

**价格**：按每 100 万令牌计费。缓存价格同时列出读取和写入成本。

**扩展思考配置**：
- **最大令牌数 4096**：Generator（生成器）使用 claude-opus-4-6，以最大推理深度处理复杂漏洞利用开发
- **最大令牌数 2048**：Coder（编码）使用 claude-sonnet-4-6，兼顾代码分析和漏洞研究
- **最大令牌数 1024**：Primary agent（主智能体）、assistant（助手）、refiner（优化）、adviser（顾问）、reflector（反思）、searcher（搜索）、installer（安装）和 pentester（渗透测试）用于针对具体任务进行集中推理
- **扩展思考**：所有 Claude 4.5+ 和 4.6 模型均支持可配置的扩展思考，可用于深度推理任务

**主要特性**：
- **扩展思考**：所有 Claude 4.5+ 和 4.6 模型都可配置思维链推理深度，用于复杂安全分析
- **自适应思考**：Claude 4.6 系列（Opus/Sonnet）会根据任务复杂度动态调整推理深度
- **提示词缓存**：读取与写入分别计价，可明显降低成本（读取为输入价格的 10%，写入为 125%）
- **扩展上下文窗口**：标准窗口为 200K 令牌；Claude Opus/Sonnet 4.6 的测试版最高支持 100 万令牌，可用于全面分析代码库
- **工具调用**：函数调用准确可靠，适合编排安全工具
- **流式输出**：为交互式渗透测试任务流实时返回响应
- **安全优先设计**：内置安全机制，帮助以负责任的方式开展安全测试
- **多模态支持**：最新模型具有视觉能力，可分析截图并评估界面安全性
- **Constitutional AI（宪法式 AI）**：通过高级安全训练提供可靠、合乎规范的安全指导

### Google AI（Gemini）提供商配置

PentAGI 通过 Google AI API 接入 Gemini 模型，支持高级多模态推理、扩展思考和上下文缓存。

#### 配置变量

| 变量                | 默认值                                      | 说明                           |
| ------------------- | ------------------------------------------- | ------------------------------ |
| `GEMINI_API_KEY`    |                                             | Google AI 服务的 API 密钥      |
| `GEMINI_SERVER_URL` | `https://generativelanguage.googleapis.com` | Google AI API 端点             |

#### 配置示例

```bash
# Gemini 基本设置
GEMINI_API_KEY=your_gemini_api_key
GEMINI_SERVER_URL=https://generativelanguage.googleapis.com

# 通过代理连接
GEMINI_API_KEY=your_gemini_api_key
PROXY_URL=http://your-proxy:8080
```

#### 支持的模型

PentAGI 支持 9 个 Gemini 模型，可使用工具调用、流式输出、思考模式和上下文缓存。标有 `*` 的模型用于默认配置。

**Gemini 3.5 系列：最新稳定版 Flash（2026 年 5 月）**

| 模型 ID                               | 思考     | 上下文  | 价格（输入/输出/缓存）     | 适用场景                                        |
| ------------------------------------- | -------- | ------- | -------------------------- | ----------------------------------------------- |
| `gemini-3.5-flash`*                   | ✅        | 1M      | $1.50/$9.00/$0.15          | 智能水平最高的 Flash 模型，在智能体和编码任务中持续保持前沿性能，搜索及基于检索结果作答的能力更强 |

**Gemini 3.1 系列：稳定版 Flash-Lite 与 Pro 预览版（2026 年 2 月至 5 月）**

| 模型 ID                               | 思考     | 上下文  | 价格（输入/输出/缓存）     | 适用场景                                        |
| ------------------------------------- | -------- | ------- | -------------------------- | ----------------------------------------------- |
| `gemini-3.1-pro-preview`*             | ✅        | 1M      | $2.00/$12.00/$0.20         | 最新旗舰模型，优化了思考过程并提高令牌效率，适合软件工程和智能体任务流 |
| `gemini-3.1-pro-preview-customtools`  | ✅        | 1M      | $2.00/$12.00/$0.20         | 自定义工具端点，优先调用 bash 和自定义工具（view_file、search_code） |
| `gemini-3.1-flash-lite`*              | ✅        | 1M      | $0.25/$1.50/$0.025         | 成本最低的稳定版多模态模型，在大规模智能体任务和低延迟应用中达到前沿水平 |

**Gemini 2.5 系列：高级思考模型（服务至 2026 年 10 月 16 日）**

| 模型 ID                                  | 思考     | 上下文  | 价格（输入/输出/缓存）     | 适用场景                                        |
| ---------------------------------------- | -------- | ------- | -------------------------- | ----------------------------------------------- |
| `gemini-2.5-pro`                         | ✅        | 1M      | $1.25/$10.00/$0.125        | 适合复杂编码和推理，也可用于高级威胁建模 |
| `gemini-2.5-flash`                       | ✅        | 1M      | $0.30/$2.50/$0.03          | 首个支持思考预算的混合推理模型，适合注重性价比的大规模评估 |
| `gemini-2.5-flash-lite`                  | ✅        | 1M      | $0.10/$0.40/$0.01          | 体量和成本最低，适合大规模使用和高吞吐扫描 |

**Gemma 4 开源模型（Apache 2.0，免费套餐）**

| 模型 ID                               | 思考     | 上下文  | 价格（输入/输出/缓存）     | 适用场景                                        |
| ------------------------------------- | -------- | ------- | -------------------------- | ----------------------------------------------- |
| `gemma-4-31b-it`                      | ✅        | 256K    | 免费/免费/免费             | 最大的开源 Gemma 4 稠密模型（约 31B 参数），支持文本和图像、140 多种语言，适合本地安全操作 |
| `gemma-4-26b-a4b-it`                  | ✅        | 256K    | 免费/免费/免费             | MoE 架构（总参数约 26B，激活参数约 3.8B），可在消费级 GPU 上高效推理，适合本地高吞吐扫描 |

**价格**：按每 100 万令牌计费（标准付费套餐）。上下文窗口指输入令牌上限。

> [!NOTE]
> **Gemini 2.5 系列停服**
>
> `gemini-2.5-pro`、`gemini-2.5-flash` 和 `gemini-2.5-flash-lite` 将于 **2026 年 10 月 16 日停服**。建议按以下方式迁移：
>
> - `gemini-2.5-pro` → `gemini-3.1-pro-preview`（同为 $2.00 输入价格档）
> - `gemini-2.5-flash` → `gemini-3.5-flash`（前沿能力有所改进）
> - `gemini-2.5-flash-lite` → `gemini-3.1-flash-lite`（输入价格同为 $0.25）

**默认模型分配（config.yml）**：
- **`gemini-3.1-pro-preview`**：`primary_agent`、`assistant`、`generator`、`refiner`、`adviser`、`coder`、`pentester`
- **`gemini-3.5-flash`**：`reflector`、`searcher`、`enricher`、`installer`
- **`gemini-3.1-flash-lite`**：`simple`、`simple_json`

**主要特性**：
- **扩展思考**：逐步推理复杂安全分析（所有 Gemini 3.x、2.5 系列及 Gemma 4 均可开关思考模式）
- **上下文缓存**：重复使用上下文时可明显降低成本（多数模型的缓存价格为输入价格的 10%）
- **超长上下文**：Gemini 聊天模型支持 100 万令牌，Gemma 4 开源模型支持 256K 令牌
- **多模态支持**：可处理文本、图像、视频、音频和 PDF，用于综合评估
- **工具调用**：通过函数调用接入 20 多种渗透测试工具
- **流式输出**：为交互式安全任务流实时返回响应
- **代码执行**：内置代码执行能力，可测试攻击性工具并验证漏洞利用
- **搜索依据**：集成 Google 搜索，用于威胁情报和 CVE 研究
- **文件搜索**：支持文档检索和 RAG，可执行基于知识库的评估
- **批处理 API**：非实时批处理的成本降低 50%
- **自定义工具端点**：专用的 `gemini-3.1-pro-preview-customtools` 路由，适合工具密集型智能体任务流，并优先使用已注册工具而不是 bash

**推理强度级别**：
- **High（高）**：以最大思考深度执行复杂的多步骤分析（`generator`）
- **Medium（中）**：兼顾推理深度与开销，适合一般智能体任务（`primary_agent`、`assistant`、`refiner`、`adviser`）
- **Low（低）**：高效思考目标明确的任务（`coder`、`installer`、`pentester`）

### AWS Bedrock 提供商配置

PentAGI 可接入 Amazon Bedrock，使用 Anthropic、Amazon、Cohere、DeepSeek、OpenAI、Qwen、Mistral 和 Moonshot 等公司提供的 20 多个基础模型。

#### 配置变量

| 变量                        | 默认值      | 说明                                                                                                |
| --------------------------- | ----------- | --------------------------------------------------------------------------------------------------- |
| `BEDROCK_REGION`            | `us-east-1` | Bedrock 服务所在的 AWS 区域                                                                         |
| `BEDROCK_DEFAULT_AUTH`      | `false`     | 使用 AWS SDK 默认凭证链（环境变量、EC2 角色、~/.aws/credentials），优先级最高                        |
| `BEDROCK_BEARER_TOKEN`      |             | Bearer 令牌认证，优先级高于静态凭证                                                                 |
| `BEDROCK_ACCESS_KEY_ID`     |             | 静态凭证的 AWS 访问密钥 ID                                                                          |
| `BEDROCK_SECRET_ACCESS_KEY` |             | 静态凭证的 AWS 秘密访问密钥                                                                         |
| `BEDROCK_SESSION_TOKEN`     |             | 临时凭证的 AWS 会话令牌（可选，与静态凭证配合使用）                                                 |
| `BEDROCK_SERVER_URL`        |             | 自定义 Bedrock 端点（用于 VPC 端点或本地测试）                                                      |

**认证优先级**：`BEDROCK_DEFAULT_AUTH` → `BEDROCK_BEARER_TOKEN` → `BEDROCK_ACCESS_KEY_ID`+`BEDROCK_SECRET_ACCESS_KEY`

#### 配置示例

```bash
# 推荐：AWS SDK 默认认证（EC2/ECS/Lambda 角色）
BEDROCK_REGION=us-east-1
BEDROCK_DEFAULT_AUTH=true

# Bearer 令牌认证（AWS STS、自定义认证）
BEDROCK_REGION=us-east-1
BEDROCK_BEARER_TOKEN=your_bearer_token

# 静态凭证（开发、测试）
BEDROCK_REGION=us-east-1
BEDROCK_ACCESS_KEY_ID=your_aws_access_key
BEDROCK_SECRET_ACCESS_KEY=your_aws_secret_key

# 使用代理和自定义端点
BEDROCK_REGION=us-east-1
BEDROCK_DEFAULT_AUTH=true
BEDROCK_SERVER_URL=https://bedrock-runtime.us-east-1.vpce-xxx.amazonaws.com
PROXY_URL=http://your-proxy:8080
```

#### 支持的模型

PentAGI 支持 21 个 AWS Bedrock 模型，可使用工具调用、流式输出和多模态能力。标有 `*` 的模型用于默认配置。

| 模型 ID                                          | 提供商          | 思考     | 多模态     | 价格（输入/输出）   | 适用场景                                |
| ------------------------------------------------ | --------------- | -------- | ---------- | -------------------- | --------------------------------------- |
| `us.amazon.nova-2-lite-v1:0`                     | Amazon Nova     | ❌        | ✅          | $0.33/$2.75          | 自适应推理，高效思考                    |
| `us.amazon.nova-premier-v1:0`                    | Amazon Nova     | ❌        | ✅          | $2.50/$12.50         | 复杂推理，高级分析                      |
| `us.amazon.nova-pro-v1:0`                        | Amazon Nova     | ❌        | ✅          | $0.80/$3.20          | 兼顾准确度、速度和成本                  |
| `us.amazon.nova-lite-v1:0`                       | Amazon Nova     | ❌        | ✅          | $0.06/$0.24          | 快速处理，大批量操作                    |
| `us.amazon.nova-micro-v1:0`                      | Amazon Nova     | ❌        | ❌          | $0.035/$0.14         | 超低延迟，实时监控                      |
| `us.anthropic.claude-opus-4-6-v1`*               | Anthropic       | ✅        | ✅          | $5.00/$25.00         | 顶尖编码能力，企业级智能体              |
| `us.anthropic.claude-sonnet-4-6`                 | Anthropic       | ✅        | ✅          | $3.00/$15.00         | 前沿智能，适合企业规模                  |
| `us.anthropic.claude-opus-4-5-20251101-v1:0`     | Anthropic       | ✅        | ✅          | $5.00/$25.00         | 持续多日的软件开发任务                  |
| `us.anthropic.claude-haiku-4-5-20251001-v1:0`*   | Anthropic       | ✅        | ✅          | $1.00/$5.00          | 性能接近前沿模型，速度快                |
| `us.anthropic.claude-sonnet-4-5-20250929-v1:0`*  | Anthropic       | ✅        | ✅          | $3.00/$15.00         | 真实场景智能体，出色的编码能力          |
| `us.anthropic.claude-sonnet-4-20250514-v1:0`     | Anthropic       | ✅        | ✅          | $3.00/$15.00         | 性能均衡，可用于生产环境                |
| `us.anthropic.claude-3-5-haiku-20241022-v1:0`    | Anthropic       | ❌        | ❌          | $0.80/$4.00          | 速度最快，扫描成本较低                  |
| `cohere.command-r-plus-v1:0`                     | Cohere          | ❌        | ❌          | $3.00/$15.00         | 大规模操作，RAG 能力较强                |
| `deepseek.v3.2`                                  | DeepSeek        | ❌        | ❌          | $0.58/$1.68          | 长上下文推理，效率较高                  |
| `openai.gpt-oss-120b-1:0`*                       | OpenAI（OSS）   | ✅        | ❌          | $0.15/$0.60          | 推理能力较强，适合科学分析              |
| `openai.gpt-oss-20b-1:0`                         | OpenAI（OSS）   | ✅        | ❌          | $0.07/$0.30          | 高效编码，软件开发                      |
| `qwen.qwen3-next-80b-a3b`                        | Qwen            | ❌        | ❌          | $0.15/$1.20          | 超长上下文，旗舰级推理                  |
| `qwen.qwen3-32b-v1:0`                            | Qwen            | ❌        | ❌          | $0.15/$0.60          | 推理能力均衡，适合研究场景              |
| `qwen.qwen3-coder-30b-a3b-v1:0`                  | Qwen            | ❌        | ❌          | $0.15/$0.60          | Vibe Coding（氛围编程），自然语言优先   |
| `qwen.qwen3-coder-next`                          | Qwen            | ❌        | ❌          | $0.45/$1.80          | 针对工具使用和函数调用优化              |
| `mistral.mistral-large-3-675b-instruct`          | Mistral         | ❌        | ✅          | $4.00/$12.00         | 高级多模态，长上下文                    |
| `moonshotai.kimi-k2.5`                           | Moonshot        | ❌        | ✅          | $0.60/$3.00          | 单一模型支持视觉、语言和代码            |

**价格**：按每 100 万令牌计费。支持思考/推理的模型在推理阶段会产生额外计算费用。

#### 已测试但不兼容的模型

部分 AWS Bedrock 模型经过测试后发现存在技术限制，因此**不受支持**：

| 模型系列                  | 不兼容原因                                                                                |
| ------------------------- | ----------------------------------------------------------------------------------------- |
| **GLM（Z.AI）**           | 工具调用格式与 Converse API 不兼容（要求字符串而不是 JSON）                              |
| **AI21 Jamba**            | 速率限制严格（每分钟 1～2 个请求），难以可靠测试或用于生产环境                           |
| **Meta Llama 3.3/3.1**    | 工具调用结果处理不稳定，会导致多轮任务流意外失败                                          |
| **Mistral Magistral**     | 模型不支持工具调用                                                                        |
| **Moonshot K2-Thinking**  | 工具调用的流式输出不稳定，不适合生产环境                                                  |
| **Qwen3-VL**              | 工具调用的流式输出不稳定，多模态与工具结合使用时会间歇性失败                              |

> [!IMPORTANT]
> **速率限制与配额管理**
>
> AWS Bedrock 为 Claude 模型提供的默认配额**非常严格**（新账号每分钟 2～20 个请求）。用于生产环境的渗透测试前，请执行以下操作：
>
> 1. 在 AWS Service Quotas 控制台为准备使用的模型**申请提高配额**
> 2. **使用 Amazon Nova 模型**，其默认配额较高，性能也较好
> 3. **启用预置吞吐量**，以稳定执行大批量测试
> 4. **监控用量**，达到配额上限时 AWS 会严格限流
>
> 如果不提高配额，任务流会频繁延迟或中断。

> [!WARNING]
> **Converse API 要求**
>
> PentAGI 通过 Amazon Bedrock **Converse API** 统一访问模型。所有受支持模型都必须具备以下能力：
>
> - ✅ 支持 Converse/ConverseStream API
> - ✅ 支持用于渗透测试任务流的工具调用（函数调用）
> - ✅ 支持流式工具调用，以便实时反馈
>
> 请在以下页面核对模型能力：[AWS Bedrock 模型功能](https://docs.aws.amazon.com/bedrock/latest/userguide/conversation-inference-supported-models-features.html)。

**主要特性**：
- **自动提示词缓存**：重复使用上下文时可降低 40%～70% 的成本（Claude 4.x 模型）
- **扩展思考**：逐步推理复杂安全分析（Claude、DeepSeek R1、OpenAI GPT）
- **多模态分析**：处理截图、图表和视频，执行综合测试（Nova、Claude、Mistral、Kimi）
- **工具调用**：通过函数调用接入 20 多种渗透测试工具
- **流式输出**：为交互式安全评估任务流实时返回响应

### DeepSeek 提供商配置

PentAGI 可接入 DeepSeek，以较低价格使用具有较强推理、编码和上下文缓存能力的模型。

#### 配置变量

| 变量                  | 默认值                     | 说明                                                |
| --------------------- | -------------------------- | --------------------------------------------------- |
| `DEEPSEEK_API_KEY`    |                            | 用于认证的 DeepSeek API 密钥                        |
| `DEEPSEEK_SERVER_URL` | `https://api.deepseek.com` | DeepSeek API 端点 URL                               |
| `DEEPSEEK_PROVIDER`   |                            | 接入 LiteLLM 时使用的提供商前缀（可选）             |

#### 配置示例

```bash
# 直连 API
DEEPSEEK_API_KEY=your_deepseek_api_key
DEEPSEEK_SERVER_URL=https://api.deepseek.com

# 通过 LiteLLM 代理连接
DEEPSEEK_API_KEY=your_litellm_key
DEEPSEEK_SERVER_URL=http://litellm-proxy:4000
DEEPSEEK_PROVIDER=deepseek  # 为 LiteLLM 添加模型名称前缀（deepseek/deepseek-v4-flash）
```

#### 支持的模型

PentAGI 支持 2 个 DeepSeek V4 模型，可使用工具调用、流式输出、思考/非思考混合模式和上下文缓存。两个模型默认均启用思考模式，也可以通过 `extra_body` 切换到非思考模式。标有 `*` 的模型用于默认配置。

| 模型 ID               | 思考     | 最大输出   | 上下文  | 价格（输入/输出/缓存）     | 适用场景                                             |
| --------------------- | -------- | ---------- | ------- | -------------------------- | ---------------------------------------------------- |
| `deepseek-v4-flash`*  | ✅ 混合   | 384K       | 1M      | $0.14/$0.28/$0.0028        | 辅助型智能体、一般对话、快速工具调用                 |
| `deepseek-v4-pro`*    | ✅ 混合   | 384K       | 1M      | $1.74/$3.48/$0.0145        | 高级推理、复杂逻辑、安全分析                         |

**价格**：按每 100 万令牌计费。缓存价格适用于从缓存读取的提示词令牌（输入缓存命中价自 2026 年 4 月 26 日起降至首发价格的 1/10）。两个模型均支持混合思考；默认开启 `thinking` 模式，传入 `extra_body.thinking.type: disabled` 可切换到非思考模式，以更快、更低成本地响应。

> **价格说明（deepseek-v4-pro）**：`deepseek-v4-pro` 的 75% 推广折扣已于 2026 年 5 月 31 日 15:59（UTC）结束。上表列出折扣结束后的标准价格。如果旧配置仍使用折扣价（$0.435/$0.87/$0.003625），请更新为当前价格，以便准确统计成本。

> DeepSeek 计划于 2026 年 7 月 24 日弃用旧模型名称 `deepseek-chat` 和
> `deepseek-reasoner`。在此之前，引用旧名称的现有用户配置仍可使用；
> 上述默认配置已改用当前 V4 名称。`deepseek-chat` 映射到
> `deepseek-v4-flash` 的非思考模式；`deepseek-reasoner` 映射到
> `deepseek-v4-flash` 的思考模式。

**默认智能体配置**：

配置策略是优先将 `deepseek-v4-flash` 用作辅助型和轻量智能体的主力模型，其输入和输出成本均为 Pro 的 1/12；`deepseek-v4-pro` 则留给复杂的多步骤推理。`installer` 智能体使用开启思考的 Flash，因为环境设置任务（shell 命令、配置修改）通常不需要 Pro 级推理。将更多智能体切换到 Pro 前，请先使用自己的工作负载进行 A/B 测试。

| 智能体角色                                  | 默认模型             | 思考     | 推理强度         | 最大输出   | 温度        | Top P |
| ------------------------------------------- | -------------------- | -------- | ---------------- | ---------- | ----------- | ----- |
| Generator（生成）/ Refiner（优化）          | `deepseek-v4-pro`    | 开启     | 高               | 32768      | （自动）    | （自动） |
| Coder（编码）                               | `deepseek-v4-pro`    | 开启     | 高               | 20480      | （自动）    | （自动） |
| Primary Agent（主智能体）/ Assistant（助手）/ Pentester（渗透测试） | `deepseek-v4-pro` | 开启 | 高 | 16384 | （自动） | （自动） |
| Adviser（导师/规划）                        | `deepseek-v4-pro`    | 开启     | 高               | 8192       | （自动）    | （自动） |
| Installer（安装）                           | `deepseek-v4-flash`  | 开启     | 高               | 12288      | （自动）    | （自动） |
| Reflector（反思）/ Searcher（搜索）/ Enricher（补充） | `deepseek-v4-flash` | 关闭 | — | 4096 | 0.5 | 0.9 |
| Simple（简单任务）/ Simple JSON             | `deepseek-v4-flash`  | 关闭     | —                | 2048       | 0.3         | 0.9   |

> **注意**：开启思考模式后，DeepSeek 会直接忽略 `temperature`、`top_p`、`presence_penalty` 和 `frequency_penalty`。设置 `reasoning_effort` 时，langchaingo 客户端会自动将 `temperature`/`top_p` 置空，因此上表将其标为“（自动）”。所有开启思考的智能体还会显式传入 `extra_body.thinking.type: enabled`，防止提供商将来更改默认值后影响现有行为。

**主要特性**：
- **混合思考模式**：通过 `extra_body.thinking.type` 在思考（深度推理）与非思考（快速响应）模式间切换
- **自动提示词缓存**：按缓存命中价格计费，可明显降低重复上下文的成本（首发价格的 1/10）
- **扩展思考**：两个 V4 模型都通过强化学习思维链处理复杂安全分析
- **较强的编码能力**：针对代码生成和漏洞利用开发进行了优化
- **长上下文**：上下文窗口为 100 万令牌，最多输出 384K 令牌
- **工具调用**：通过函数调用接入 20 多种渗透测试工具
- **流式输出**：为交互式任务流实时返回响应
- **多语言**：对中文和英文的支持较好
- **其他功能**：JSON 输出、聊天前缀补全（测试版）、FIM/中间填充补全（仅限非思考模式）

**并发限制**：`deepseek-v4-flash` 支持 2500 个并发请求；`deepseek-v4-pro` 支持 500 个并发请求。

**LiteLLM 集成**：通过 LiteLLM 代理使用 PentAGI 默认配置时，设置 `DEEPSEEK_PROVIDER=deepseek` 可为模型名称添加前缀；直连 API 时请留空。

### GLM 提供商配置

PentAGI 可接入智谱 AI（Z.AI）的 GLM，使用由清华大学研发、采用 MoE 架构并具有较强推理和智能体能力的高级语言模型。

#### 配置变量

| 变量              | 默认值                          | 说明                                                       |
| ----------------- | ------------------------------- | ---------------------------------------------------------- |
| `GLM_API_KEY`     |                                 | 用于认证的 GLM API 密钥                                    |
| `GLM_SERVER_URL`  | `https://api.z.ai/api/paas/v4`  | GLM API 端点 URL（国际站）                                 |
| `GLM_PROVIDER`    |                                 | 接入 LiteLLM 时使用的提供商前缀（可选）                    |

#### 配置示例

```bash
# 直连 API（国际站端点）
GLM_API_KEY=your_glm_api_key
GLM_SERVER_URL=https://api.z.ai/api/paas/v4

# 其他端点
GLM_SERVER_URL=https://open.bigmodel.cn/api/paas/v4  # 中国站
GLM_SERVER_URL=https://api.z.ai/api/coding/paas/v4   # 编码专用

# 通过 LiteLLM 代理连接
GLM_API_KEY=your_litellm_key
GLM_SERVER_URL=http://litellm-proxy:4000
GLM_PROVIDER=zai  # 为 LiteLLM 添加模型名称前缀（zai/glm-4）
```

#### 支持的模型

PentAGI 支持 13 个 GLM 模型，可使用工具调用、流式输出、混合思考模式和提示词缓存。标有 `*` 的模型用于默认配置。思考模式通过 `extra_body.thinking.type`（`"enabled"`/`"disabled"`）控制；与 Kimi 不同，GLM 在两种模式下对 temperature 的限制都较少。

**GLM-5.x 系列：最新一代（200K 上下文，最大输出 128K）**

| 模型 ID          | 思考     | 上下文  | 最大输出   | 价格（输入/输出/缓存）     | 适用场景                                                            |
| ---------------- | -------- | ------- | ---------- | -------------------------- | ------------------------------------------------------------------- |
| `glm-5.1`*       | ✅ 混合   | 200K    | 128K       | $1.40/$4.40/$0.26          | 最新旗舰：支持持续 8 小时自主执行，对标 Claude Opus 4.6（`generator`/`refiner`/`adviser`/`coder`/`pentester` 的默认模型） |
| `glm-5`          | ✅ 混合   | 200K    | 128K       | $1.00/$3.20/$0.20          | 面向智能体工程的基础模型，MoE 总参数 744B/激活参数 40B，编码能力达到 Claude Opus 4.5 水平 |
| `glm-5-turbo`*   | ✅ 混合   | 200K    | 128K       | $1.20/$4.00/$0.24          | OpenClaw 原生模型，针对工具调用、持久任务和长链执行优化（`primary_agent`/`assistant` 的默认模型） |

**GLM-4.7 系列：支持交错思考的高端模型**

| 模型 ID           | 思考     | 上下文  | 最大输出   | 价格（输入/输出/缓存）     | 适用场景                                            |
| ----------------- | -------- | ------- | ---------- | -------------------------- | --------------------------------------------------- |
| `glm-4.7`         | ✅ 混合   | 200K    | 128K       | $0.60/$2.20/$0.11          | 增强编程能力，多步骤推理稳定                        |
| `glm-4.7-flashx`  | ✅ 混合   | 200K    | 128K       | $0.07/$0.40/$0.01          | 价格很低并使用优先 GPU，但 RPM 限制较低，不适合高频使用 |
| `glm-4.7-flash`   | ✅ 混合   | 200K    | 128K       | 免费/免费/免费             | 免费的约 30B 参数先进模型，仅支持 1 个并发请求      |

**GLM-4.6 系列：均衡的自动思考模型**

| 模型 ID   | 思考     | 上下文  | 最大输出   | 价格（输入/输出/缓存）     | 适用场景                                          |
| --------- | -------- | ------- | ---------- | -------------------------- | ------------------------------------------------- |
| `glm-4.6` | ✅ 自动   | 200K    | 128K       | $0.60/$2.20/$0.11          | 性能均衡，支持流式工具调用，令牌效率较高          |

**GLM-4.5 系列：统一支持推理、编码和智能体**

| 模型 ID          | 思考     | 上下文  | 最大输出   | 价格（输入/输出/缓存）     | 适用场景                                          |
| ---------------- | -------- | ------- | ---------- | -------------------------- | ------------------------------------------------- |
| `glm-4.5`        | ✅ 自动   | 128K    | 96K        | $0.60/$2.20/$0.11          | 统一模型，MoE 总参数 355B/激活参数 32B            |
| `glm-4.5-x`      | ✅ 自动   | 128K    | 96K        | $2.20/$8.90/$0.45          | 超高速高端模型，延迟最低                          |
| `glm-4.5-air`*   | ✅ 自动   | 128K    | 96K        | $0.20/$1.10/$0.03          | 高性价比 MoE，总参数 106B/激活参数 12B（`simple`/`simple_json`/`reflector`/`searcher`/`enricher`/`installer` 的默认模型） |
| `glm-4.5-airx`   | ✅ 自动   | 128K    | 96K        | $1.10/$4.50/$0.22          | 使用优先 GPU 加速的 Air 模型                      |
| `glm-4.5-flash`  | ✅ 自动   | 128K    | 96K        | 免费/免费/免费             | 免费，支持推理、编码和智能体任务                  |

**旧版 GLM-4：稠密架构**

| 模型 ID               | 思考     | 上下文  | 最大输出   | 价格（输入/输出）   | 适用场景                                      |
| --------------------- | -------- | ------- | ---------- | -------------------- | --------------------------------------------- |
| `glm-4-32b-0414-128k` | ❌        | 128K    | 16K        | $0.10/$0.10          | 超低成本的稠密 32B 模型，用于无需推理的解析任务 |

**价格**：按每 100 万令牌计费。缓存价格适用于提示词缓存命中；根据 Z.AI 的推广活动，缓存存储目前免费。GLM-4-32B 不支持缓存。

**默认智能体配置**：

配置策略是让 `glm-5.1`（最新旗舰，输入价格 $1.40）处理关键推理，`glm-5-turbo`（OpenClaw 原生，针对智能体优化）负责编排，`glm-4.5-air`（低成本 MoE，支持混合思考且 RPM 稳定）用于所有辅助型和安装智能体。`glm-4.7-flashx` 的 RPM（每分钟请求数）限制较低，高频调用时容易出现 429 错误，因此不作为默认模型。

| 智能体角色                          | 默认模型      | 思考     | 温度        | Top P | 最大输出   |
| ----------------------------------- | ------------- | -------- | ----------- | ----- | ---------- |
| Generator（生成）/ Refiner（优化）  | `glm-5.1`     | 开启     | 1.0         | 0.95  | 32768      |
| Coder（编码）                       | `glm-5.1`     | 开启     | 1.0         | 0.95  | 20480      |
| Adviser（顾问）/ Pentester（渗透测试） | `glm-5.1`  | 开启     | 1.0         | 0.95  | 16384      |
| Primary Agent（主智能体）/ Assistant（助手） | `glm-5-turbo` | 开启 | 1.0 | 0.95 | 16384 |
| Installer（安装）                   | `glm-4.5-air` | 开启     | 1.0         | 0.95  | 16384      |
| Simple（简单任务）/ Reflector（反思） | `glm-4.5-air` | 关闭   | 0.6         | 0.9   | 8192       |
| Searcher（搜索）/ Enricher（补充）/ Simple JSON | `glm-4.5-air` | 关闭 | 0.6 | 0.9 | 4096 |

> **温度说明**：根据 Z.AI 文档，GLM 在思考和非思考模式下都接受 `1.0` 与 `0.6`。langchaingo 的 `IsReasoningModel` 会匹配 `glm-4.5*`/`glm-4.6*`/`glm-4.7*` 前缀，并在 `createChatRequest` 中强制将温度覆盖为 1.0。这对 GLM 没有影响（Kimi 则不同），但意味着 YAML 中这些模型的温度值仅供参考。`glm-5`/`glm-5.1`/`glm-5-turbo` 不会被匹配，因此显式设置的值会原样传入。

**思考模式**：
- **混合模式**（GLM-5.x、GLM-4.7）：通过 `extra_body.thinking.type` 显式开关
- **自动模式**（GLM-4.6、GLM-4.5 系列）：模型自动判断何时需要推理
- **保留思考**（Z.AI Coding 功能）：PentAGI 中所有开启思考的智能体还会传入 `extra_body.thinking.clear_thinking: false`，在整个对话中保留助手前几轮的 `reasoning_content`。标准 API 端点（`/api/paas/v4`）要求显式设置此项；Coding Plan 端点则默认开启。这样可以保持多轮工具调用链中的推理连续性，并提高缓存命中率
- 所有开启思考的智能体还会传入 `extra_body.tool_choice: auto`，防止提供商默认行为变化后影响工具选择

**主要特性**：
- **长周期任务**：GLM-5.1 支持持续 8 小时自主执行，适合复杂的多阶段智能体任务流
- **OpenClaw 原生编排**：GLM-5-Turbo 专门针对工具调用、指令遵循和长链执行优化
- **提示词缓存**：重复使用上下文时可明显降低成本（表中已列出缓存输入价格）
- **超长上下文**：GLM-5.x/4.7/4.6 系列支持 200K 令牌
- **MoE 架构**：GLM-5/5.1 为总参数 744B/激活参数 40B，GLM-4.5 为 355B/32B，GLM-4.5-Air 为 106B/12B
- **工具调用**：通过函数调用接入 20 多种渗透测试工具
- **流式输出**：支持实时流式输出及流式工具调用（GLM-4.6+）
- **多语言**：具有出色的中文和英文自然语言处理能力
- **免费选项**：GLM-4.7-Flash 和 GLM-4.5-Flash 可用于原型开发和实验

**LiteLLM 集成**：通过 LiteLLM 代理使用 PentAGI 默认配置时，设置 `GLM_PROVIDER=zai` 可为模型名称添加前缀；直连 API 时请留空。

### Kimi 提供商配置

PentAGI 可接入 Moonshot AI 的 Kimi，使用具有多模态能力的超长上下文模型，适合分析大型代码库和文档。

#### 配置变量

| 变量               | 默认值                       | 说明                                                |
| ------------------ | ---------------------------- | --------------------------------------------------- |
| `KIMI_API_KEY`     |                              | 用于认证的 Kimi API 密钥                            |
| `KIMI_SERVER_URL`  | `https://api.moonshot.ai/v1` | Kimi API 端点 URL（国际站）                         |
| `KIMI_PROVIDER`    |                              | 接入 LiteLLM 时使用的提供商前缀（可选）             |

#### 配置示例

```bash
# 直连 API（国际站端点）
KIMI_API_KEY=your_kimi_api_key
KIMI_SERVER_URL=https://api.moonshot.ai/v1

# 其他端点
KIMI_SERVER_URL=https://api.moonshot.cn/v1  # 中国站

# 通过 LiteLLM 代理连接
KIMI_API_KEY=your_litellm_key
KIMI_SERVER_URL=http://litellm-proxy:4000
KIMI_PROVIDER=moonshot  # 为 LiteLLM 添加模型名称前缀（moonshot/kimi-k2.5）
```

#### 支持的模型

PentAGI 支持 8 个 Kimi/Moonshot 模型，可使用工具调用、流式输出、混合思考模式和多模态能力（K2.x 支持文本/图像/视频）。Moonshot 已于 2026 年 5 月 25 日弃用所有旧版 `kimi-k2-*` 模型（turbo-preview、0905-preview、0711-preview、thinking、thinking-turbo），此处**未将其列入**。标有 `*` 的模型用于默认配置。

**Kimi K2.x 系列：多模态旗舰模型**

| 模型 ID          | 思考     | 多模态     | 上下文  | 价格（输入未命中/输出/缓存命中）        | 适用场景                                                |
| ---------------- | -------- | ---------- | ------- | --------------------------------------- | ------------------------------------------------------- |
| `kimi-k2.6`*     | ✅ 混合   | ✅          | 256K    | $0.95 / $4.00 / $0.16                   | 最新旗舰：原生多模态，编码和指令遵循能力更强（`generator`/`refiner`/`adviser`/`coder`/`pentester` 的默认模型） |
| `kimi-k2.5`*     | ✅ 混合   | ✅          | 256K    | $0.60 / $3.00 / $0.10                   | 上一代同架构模型，输入成本低 36%（`primary`/`assistant`/`installer`/辅助型智能体的默认模型） |

**Moonshot V1 系列：参数灵活的生成模型**

| 模型 ID             | 思考     | 多模态     | 上下文  | 价格（输入/输出）     | 适用场景                                       |
| ------------------- | -------- | ---------- | ------- | ---------------------- | ---------------------------------------------- |
| `moonshot-v1-8k`    | ❌        | ❌          | 8K      | $0.20 / $2.00          | 短文本生成，成本很低                           |
| `moonshot-v1-32k`   | ❌        | ❌          | 32K     | $1.00 / $3.00          | 长文本生成                                     |
| `moonshot-v1-128k`  | ❌        | ❌          | 128K    | $2.00 / $5.00          | 超长上下文                                     |

**Moonshot V1 Vision 系列：图像理解**

| 模型 ID                           | 思考     | 多模态     | 上下文  | 价格（输入/输出）     | 适用场景                                |
| --------------------------------- | -------- | ---------- | ------- | ---------------------- | --------------------------------------- |
| `moonshot-v1-8k-vision-preview`   | ❌        | ✅          | 8K      | $0.20 / $2.00          | 视觉与短上下文                          |
| `moonshot-v1-32k-vision-preview`  | ❌        | ✅          | 32K     | $1.00 / $3.00          | 视觉与中等长度上下文                    |
| `moonshot-v1-128k-vision-preview` | ❌        | ✅          | 128K    | $2.00 / $5.00          | 视觉与长上下文                          |

**价格**：按每 100 万令牌计费。缓存价格适用于从自动上下文缓存读取的提示词令牌（仅 Kimi K2.x 模型支持缓存）。

> **关键：Kimi K2.6/K2.5 参数限制**：任何不符合以下要求的值都会使 API 返回 `invalid_request_error`：
> - `temperature`：思考模式下必须为 `1.0`，非思考模式下必须为 `0.6`
> - `top_p`：必须为 `0.95`
> - `n`：必须为 `1`
> - `presence_penalty` 和 `frequency_penalty`：必须为 `0`（不可修改）
>
> Moonshot V1 模型使用兼容 OpenAI 的标准参数，没有上述限制。

**默认智能体配置**：

配置策略是优先将 `kimi-k2.5` 用作性价比较高的主力模型（输入成本比 `kimi-k2.6` 低 36%），`kimi-k2.6` 则留给关键推理。所有 `kimi-k2.x` 智能体都使用 API 要求的固定参数（temp/top_p/n），并显式设置 `extra_body.thinking.type`。对于开启思考的智能体，还会设置 `extra_body.thinking.keep: "all"`，在多轮工具调用链中保留历史 `reasoning_content`；否则 Moonshot 会返回 `thinking is enabled but reasoning_content is missing`。

| 智能体角色                                   | 默认模型      | 思考     | 温度        | Top P | 最大输出   |
| -------------------------------------------- | ------------- | -------- | ----------- | ----- | ---------- |
| Generator（生成）/ Refiner（优化）           | `kimi-k2.6`   | 开启（keep=all）   | 1.0 | 0.95 | 32768 |
| Coder（编码）                                | `kimi-k2.6`   | 开启（keep=all）   | 1.0 | 0.95 | 20480 |
| Pentester（渗透测试）                        | `kimi-k2.6`   | 开启（keep=all）   | 1.0 | 0.95 | 16384 |
| Adviser（导师/规划）                         | `kimi-k2.6`   | 开启（keep=all）   | 1.0 | 0.95 | 8192  |
| Primary Agent（主智能体）/ Assistant（助手） | `kimi-k2.5`   | 开启（keep=all）   | 1.0 | 0.95 | 16384 |
| Installer（安装）                            | `kimi-k2.5`   | 开启（keep=all）   | 1.0 | 0.95 | 12288 |
| Reflector（反思）/ Searcher（搜索）/ Enricher（补充） | `kimi-k2.5` | 关闭 | 0.6 | 0.95 | 4096 |
| Simple（简单任务）/ Simple JSON              | `kimi-k2.5`   | 关闭               | 0.6 | 0.95 | 2048  |

**主要特性**：
- **超长上下文**：K2.x 最高支持 256K 令牌，可全面分析代码库和文档
- **原生多模态**：K2.6/K2.5 原生支持文本、图像和视频输入
- **混合思考**：K2.6/K2.5 通过 `extra_body.thinking.type` 在思考与非思考模式间切换
- **保留思考**（K2.6）：`thinking.keep: "all"` 会跨轮保留历史 `reasoning_content`，这是多轮工具调用链的必要设置
- **自动上下文缓存**：K2.x 模型会缓存重复前缀（K2.6 和 K2.5 的缓存价格均约为未命中价格的 17%）
- **工具调用**：K2.x 和 Moonshot V1 完整支持函数调用
- **自我纠正**：K2.6 改进了指令遵循和自我纠正能力
- **多语言**：对中文、英文及其他多种语言的支持较好

**思考与工具调用并用的多轮对话**：PentAGI 的通用推理保留模式（`TextPartWithReasoning` + `WithPreserveReasoningContent`）会自动按要求的 TextContent → ToolCall 顺序回传 `reasoning_content`，从而避免 Moonshot 报出 `thinking is enabled but reasoning_content is missing in assistant tool call message` 错误。

**LiteLLM 集成**：通过 LiteLLM 代理使用 PentAGI 默认配置时，设置 `KIMI_PROVIDER=moonshot` 可为模型名称添加前缀；直连 API 时请留空。

### Qwen 提供商配置

PentAGI 可接入阿里云百炼（DashScope）的 Qwen，使用支持推理与上下文缓存的多语言模型。

#### 配置变量

| 变量               | 默认值                                                 | 说明                                                |
| ------------------ | ------------------------------------------------------ | --------------------------------------------------- |
| `QWEN_API_KEY`     |                                                        | 用于认证的 Qwen API 密钥                            |
| `QWEN_SERVER_URL`  | `https://dashscope-us.aliyuncs.com/compatible-mode/v1` | Qwen API 端点 URL（国际站）                         |
| `QWEN_PROVIDER`    |                                                        | 接入 LiteLLM 时使用的提供商前缀（可选）             |

#### 配置示例

```bash
# 直连 API（全球/美国端点）
QWEN_API_KEY=your_qwen_api_key
QWEN_SERVER_URL=https://dashscope-us.aliyuncs.com/compatible-mode/v1

# 其他端点
QWEN_SERVER_URL=https://dashscope-intl.aliyuncs.com/compatible-mode/v1  # 国际站（新加坡）
QWEN_SERVER_URL=https://dashscope.aliyuncs.com/compatible-mode/v1       # 中国大陆（北京）

# 通过 LiteLLM 代理连接
QWEN_API_KEY=your_litellm_key
QWEN_SERVER_URL=http://litellm-proxy:4000
QWEN_PROVIDER=dashscope  # 为 LiteLLM 添加模型名称前缀（dashscope/qwen-plus）
```

#### 支持的模型

PentAGI 精选了 33 个适合智能体任务流的 Qwen 模型，涵盖文本推理、代码生成和视觉语言（浏览器截图）任务。所有模型均使用非快照主别名，支持工具调用、流式输出、思考模式和上下文缓存。标有 `*` 的模型用于默认配置。

**旗舰模型（顶级推理）**

| 模型 ID                      | 思考     | 国际站 | 全球/美国 | 中国 | 价格（输入/输出/缓存）     | 适用场景                                                |
| ---------------------------- | -------- | ---- | --------- | ----- | -------------------------- | ------------------------------------------------------- |
| `qwen3.7-max`*               | ✅        | ✅    | ✅         | ✅     | $2.50/$7.50/$0.50          | 面向智能体时代的下一代旗舰模型（`generator`/`refiner`/`adviser` 的默认模型） |
| `qwen3.6-max-preview`        | ✅        | ✅    | ✅         | ✅     | $1.30/$7.80/$0.13          | Max 预览版，增强了 Vibe Coding（氛围编程）和前端开发能力 |
| `qwen3-max`                  | ✅        | ✅    | ✅         | ✅     | $1.20/$6.00/$0.24          | 上一代旗舰模型，改进了智能体编程能力                    |
| `qwen-plus`                  | ✅        | ✅    | ✅         | ✅     | $0.40/$4.00/$0.08          | 基于 Qwen3 的 Plus 模型，可切换思考模式                 |

**均衡型 Plus 模型（中端）**

| 模型 ID                      | 思考     | 国际站 | 全球/美国 | 中国 | 价格（输入/输出/缓存）     | 适用场景                                                |
| ---------------------------- | -------- | ---- | --------- | ----- | -------------------------- | ------------------------------------------------------- |
| `qwen3.6-plus`*              | ✅        | ✅    | ✅         | ✅     | $0.50/$3.00/$0.05          | 原生视觉语言 Plus 模型，支持智能体编码（`primary`/`assistant`/`pentester` 的默认模型） |
| `qwen3.5-plus`               | ✅        | ✅    | ✅         | ✅     | $0.40/$2.40/$0.04          | 上一代原生视觉语言模型，多模态能力较强                  |

**快速 Flash 模型（成本优化）**

| 模型 ID                      | 思考     | 国际站 | 全球/美国 | 中国 | 价格（输入/输出/缓存）     | 适用场景                                                |
| ---------------------------- | -------- | ---- | --------- | ----- | -------------------------- | ------------------------------------------------------- |
| `qwen3.6-flash`              | ✅        | ✅    | ✅         | ✅     | $0.25/$1.50/$0.025         | 最新 Flash 模型，智能体编码能力明显增强                |
| `qwen3.5-flash`*             | ✅        | ✅    | ✅         | ✅     | $0.10/$0.40/$0.01          | 超高速轻量模型（`simple`/`reflector`/`searcher`/`enricher` 的默认模型） |
| `qwen-flash`                 | ✅        | ✅    | ✅         | ✅     | $0.05/$0.40/$0.01          | Qwen3 系列 Flash 模型，支持 100 万令牌上下文和分级定价 |

**代码专用模型**

| 模型 ID                      | 思考     | 国际站 | 全球/美国 | 中国 | 价格（输入/输出/缓存）     | 适用场景                                                |
| ---------------------------- | -------- | ---- | --------- | ----- | -------------------------- | ------------------------------------------------------- |
| `qwen3-coder-plus`*          | ❌        | ✅    | ✅         | ✅     | $1.00/$5.00/$0.20          | 编码能力较强的智能体模型，支持自主编程（`coder` 的默认模型） |
| `qwen3-coder-flash`*         | ❌        | ✅    | ✅         | ✅     | $0.30/$1.50/$0.06          | 快速生成代码，多轮工具调用稳定（`installer` 的默认模型） |
| `qwen3-coder-next`           | ❌        | ✅    | ✅         | ✅     | $0.30/$1.50/—              | 开源代码生成模型，同等规模下达到先进水平               |

**视觉语言模型（浏览器和截图分析）**

| 模型 ID                      | 思考     | 国际站 | 全球/美国 | 中国 | 价格（输入/输出/缓存）     | 适用场景                                                |
| ---------------------------- | -------- | ---- | --------- | ----- | -------------------------- | ------------------------------------------------------- |
| `qwen3-vl-plus`              | ✅        | ✅    | ✅         | ✅     | $0.20/$1.60/$0.04          | 具有视觉智能体能力的 VL 模型，支持超长视频理解         |
| `qwen3-vl-flash`             | ✅        | ✅    | ✅         | ✅     | $0.05/$0.40/$0.01          | 小型 VL 模型，支持 2D/3D 定位，可用于浏览器截图初筛    |
| `qvq-max`                    | ✅        | ✅    | ✅         | ✅     | $1.20/$4.80/—              | 支持思维链的视觉推理                                    |

**开源 Qwen3.6 系列**

| 模型 ID                      | 思考     | 国际站 | 全球/美国 | 中国 | 价格（输入/输出/缓存）     | 适用场景                                                |
| ---------------------------- | -------- | ---- | --------- | ----- | -------------------------- | ------------------------------------------------------- |
| `qwen3.6-27b`                | ✅        | ✅    | ✅         | ✅     | $0.60/$3.60/—              | 混合架构的原生 VL 模型，可在本地部署                   |
| `qwen3.6-35b-a3b`            | ✅        | ✅    | ✅         | ✅     | $0.25/$1.49/—              | 高效的 35B MoE（约 3B 激活参数），适合持续监控         |

**开源 Qwen3.5 系列**

| 模型 ID                      | 思考     | 国际站 | 全球/美国 | 中国 | 价格（输入/输出/缓存）     | 适用场景                                                |
| ---------------------------- | -------- | ---- | --------- | ----- | -------------------------- | ------------------------------------------------------- |
| `qwen3.5-397b-a17b`          | ✅        | ✅    | ✅         | ✅     | $0.60/$3.60/—              | 最大的 397B 参数模型（约 17B 激活参数），推理能力出色  |
| `qwen3.5-122b-a10b`          | ✅        | ✅    | ✅         | ✅     | $0.40/$3.20/—              | 大型 122B 参数模型（约 10B 激活参数），各项能力较均衡  |
| `qwen3.5-35b-a3b`            | ✅        | ✅    | ✅         | ✅     | $0.25/$2.00/—              | 高效的 35B MoE（约 3B 激活参数），性价比较高           |
| `qwen3.5-27b`                | ✅        | ✅    | ✅         | ✅     | $0.30/$2.40/—              | 中型 27B 模型，采用混合线性注意力与稀疏 MoE            |

**开源 Qwen3 Coder 系列**

| 模型 ID                               | 思考     | 国际站 | 全球/美国 | 中国 | 价格（输入/输出/缓存）     | 适用场景                                                |
| ------------------------------------- | -------- | ---- | --------- | ----- | -------------------------- | ------------------------------------------------------- |
| `qwen3-coder-480b-a35b-instruct`      | ❌        | ✅    | ✅         | ✅     | $1.50/$7.50/—              | 最大的开源编码 MoE（480B 总参数/约 35B 激活参数）      |
| `qwen3-coder-30b-a3b-instruct`        | ❌        | ✅    | ✅         | ✅     | $0.45/$2.25/—              | 高效的 30B MoE（约 3B 激活参数），可处理仓库级任务     |

**开源 Qwen3 稠密与 MoE 系列**

| 模型 ID                               | 思考     | 国际站 | 全球/美国 | 中国 | 价格（输入/输出/缓存）     | 适用场景                                                |
| ------------------------------------- | -------- | ---- | --------- | ----- | -------------------------- | ------------------------------------------------------- |
| `qwen3-next-80b-a3b-thinking`         | ✅        | ✅    | ✅         | ✅     | $0.15/$1.20/—              | 新一代 80B MoE（约 3B 激活参数），仅支持思考模式       |
| `qwen3-next-80b-a3b-instruct`         | ❌        | ✅    | ✅         | ✅     | $0.15/$1.20/—              | 新一代 80B MoE 指令遵循模型                            |
| `qwen3-235b-a22b`                     | ✅        | ✅    | ✅         | ✅     | $0.70/$8.40/—              | 双模式 235B MoE（约 22B 激活参数）                     |
| `qwen3-32b`                           | ✅        | ✅    | ✅         | ✅     | $0.16/$0.64/—              | 用途广泛的 32B 稠密双模式模型                          |
| `qwen3-30b-a3b`                       | ✅        | ✅    | ✅         | ✅     | $0.20/$2.40/—              | 高效的 30B MoE（约 3B 激活参数）                       |
| `qwen3-14b`                           | ✅        | ✅    | ✅         | ✅     | $0.35/$4.20/—              | 中型 14B 稠密模型，兼顾性能和成本                      |
| `qwen3-8b`                            | ✅        | ✅    | ✅         | ✅     | $0.18/$2.10/—              | 小型 8B 稠密模型，效率较高                             |
| `qwen3-4b`                            | ✅        | ✅    | ✅         | ✅     | $0.11/$1.26/—              | 轻量 4B 稠密模型，适合简单任务                         |
| `qwen3-1.7b`                          | ✅        | ✅    | ✅         | ✅     | $0.11/$1.26/—              | 超小型 1.7B 模型，适合基本检查                         |
| `qwen3-0.6b`                          | ✅        | ✅    | ✅         | ✅     | $0.11/$1.26/—              | 最小的 0.6B 模型，适合边缘监控                         |

**价格**：按每 100 万令牌计费。缓存价格指隐式缓存命中时的价格（如支持）；开源 MoE/稠密模型未提供缓存价格。分级定价模型（Max/Plus）列出最低档价格（输入通常不超过 32K 或 256K）；根据阿里云定价，使用更大上下文时费率更高。

**区域可用性**：
- **Intl（国际站）**：新加坡区域（`dashscope-intl.aliyuncs.com`）
- **Global/US（全球/美国）**：美国弗吉尼亚区域（`dashscope-us.aliyuncs.com`）
- **China（中国）**：中国大陆北京区域（`dashscope.aliyuncs.com`）

**默认智能体配置**：
| 智能体角色                                       | 默认模型             | 档位      |
| ------------------------------------------------ | -------------------- | --------- |
| Generator（生成）/ Refiner（优化）/ Adviser（规划、导师） | `qwen3.7-max` | 旗舰 |
| Primary（主智能体）/ Assistant（助手）/ Pentester（渗透测试） | `qwen3.6-plus` | 均衡 |
| Coder（漏洞利用开发）                            | `qwen3-coder-plus`   | Code+     |
| Installer（环境设置）                            | `qwen3-coder-flash`  | 快速编码  |
| Simple（简单任务）/ Reflector（反思）/ Searcher（搜索）/ Enricher（补充） | `qwen3.5-flash` | 快速 |

**主要特性**：
- **面向智能体设计**：Qwen3.7-Max 专为长周期自主执行和工具调用而设计
- **自动上下文缓存**：通过隐式缓存将重复上下文的成本降低 30%～50%
- **扩展思考**：通过思维链推理复杂安全分析（Qwen3.7/3.6/3.5/3-Max、QVQ-Max）
- **代码专用模型**：Qwen3-Coder 系列支持多轮工具交互和仓库级理解
- **视觉语言**：Qwen3-VL 系列可对浏览器截图进行初步分析、2D/3D 定位和 OCR 级分析
- **工具调用**：通过函数调用接入 20 多种渗透测试工具
- **流式输出**：为交互式任务流实时返回响应
- **多语言**：对中文、英文及其他多种语言的支持较好
- **开源变体**：提供 0.6B 至 480B 的稠密与 MoE 模型，可在本地或隔离网络中部署

**LiteLLM 集成**：通过 LiteLLM 代理使用 PentAGI 默认配置时，设置 `QWEN_PROVIDER=dashscope` 可为模型名称添加前缀；直连 API 时请留空。

#### 其他集成方式

DashScope 完全兼容 OpenAI 接口，因此还可以通过标准 OpenAI 客户端，将 Qwen 用于 PentAGI 的另外两个子系统。

**用作嵌入模型提供商**（`text-embedding-v4`，参见[阿里云百炼定价](https://modelstudio.console.alibabacloud.com/ap-southeast-1?tab=doc#/doc/?type=model&url=prices)）：

```bash
EMBEDDING_PROVIDER=openai
EMBEDDING_URL=https://dashscope-intl.aliyuncs.com/compatible-mode/v1  # 国际站（新加坡）
# EMBEDDING_URL=https://dashscope.aliyuncs.com/compatible-mode/v1     # 中国大陆
EMBEDDING_KEY=sk-*******
EMBEDDING_MODEL=text-embedding-v4
EMBEDDING_BATCH_SIZE=         # 可选，使用默认值
EMBEDDING_STRIP_NEW_LINES=    # 可选，使用默认值
```

> 注意：DashScope 全球/美国端点（`dashscope-us.aliyuncs.com`）**不提供**嵌入 API。使用 `text-embedding-v4` 时，请选择国际站或中国站端点。

**用作 OpenAI 兼容的自定义 LLM 提供商**：除了专用的 `QWEN_*` 变量，还可以通过 PentAGI 的 OpenAI 兼容自定义提供商接入任意 Qwen 聊天模型。将 `OPENAI_SERVER_URL`（或自定义提供商条目）指向 DashScope 的 `/compatible-mode/v1` 端点，再选择所需的 Qwen 模型名称即可。如果已经通过一个兼容 OpenAI 接口的客户端统一管理所有模型流量（例如与 LiteLLM/OneAPI 代理共用），这种方式会更方便。

## 高级设置

### 集成 Langfuse

Langfuse 可用于监控并分析 AI 智能体的运行情况。

1. 在现有 `.env` 文件中配置 Langfuse 环境变量。

<details>
    <summary>Langfuse 重要环境变量</summary>

### 数据库凭证
- `LANGFUSE_POSTGRES_USER` 和 `LANGFUSE_POSTGRES_PASSWORD`：Langfuse PostgreSQL 凭证
- `LANGFUSE_CLICKHOUSE_USER` 和 `LANGFUSE_CLICKHOUSE_PASSWORD`：ClickHouse 凭证
- `LANGFUSE_REDIS_AUTH`：Redis 密码

### 加密与安全密钥
- `LANGFUSE_SALT`：Langfuse 网页界面进行哈希处理时使用的盐值
- `LANGFUSE_ENCRYPTION_KEY`：加密密钥（以十六进制表示的 32 字节数据）
- `LANGFUSE_NEXTAUTH_SECRET`：NextAuth 密钥

### 管理员凭证
- `LANGFUSE_INIT_USER_EMAIL`：管理员邮箱
- `LANGFUSE_INIT_USER_PASSWORD`：管理员密码
- `LANGFUSE_INIT_USER_NAME`：管理员用户名

### API 密钥与令牌
- `LANGFUSE_INIT_PROJECT_PUBLIC_KEY`：项目公钥（PentAGI 端也会使用）
- `LANGFUSE_INIT_PROJECT_SECRET_KEY`：项目密钥（PentAGI 端也会使用）

### S3 存储
- `LANGFUSE_S3_ACCESS_KEY_ID`：S3 访问密钥 ID
- `LANGFUSE_S3_SECRET_ACCESS_KEY`：S3 秘密访问密钥

</details>

2. 在 `.env` 文件中为 PentAGI 服务启用 Langfuse 集成。

```bash
LANGFUSE_BASE_URL=http://langfuse-web:3000
LANGFUSE_PROJECT_ID= # 默认值：${LANGFUSE_INIT_PROJECT_ID} 的值
LANGFUSE_PUBLIC_KEY= # 默认值：${LANGFUSE_INIT_PROJECT_PUBLIC_KEY} 的值
LANGFUSE_SECRET_KEY= # 默认值：${LANGFUSE_INIT_PROJECT_SECRET_KEY} 的值
```

3. 启动 Langfuse 服务栈：

```bash
curl -O https://raw.githubusercontent.com/vxcontrol/pentagi/master/docker-compose-langfuse.yml
docker compose -f docker-compose.yml -f docker-compose-langfuse.yml up -d
```

打开 [localhost:4000](http://localhost:4000)，使用 `.env` 文件中的凭证登录 Langfuse 网页界面：

- `LANGFUSE_INIT_USER_EMAIL`：管理员邮箱
- `LANGFUSE_INIT_USER_PASSWORD`：管理员密码

### 监控与可观测性

如需详细跟踪系统运行情况，可以接入监控工具。

1. 在 `.env` 文件中为 PentAGI 启用 OpenTelemetry 及全部可观测性服务。

```bash
OTEL_HOST=otelcol:8148
```

2. 启动可观测性服务栈：

```bash
curl -O https://raw.githubusercontent.com/vxcontrol/pentagi/master/docker-compose-observability.yml
docker compose -f docker-compose.yml -f docker-compose-observability.yml up -d
```

打开 [localhost:3000](http://localhost:3000) 访问 Grafana 网页界面。

> [!NOTE]
> 如需同时使用可观测性服务栈与 Langfuse，请在 `.env` 文件中启用集成，将 `LANGFUSE_OTEL_EXPORTER_OTLP_ENDPOINT` 设为 `http://otelcol:4318`。
>
> 如需同时启动所有可用服务栈（Langfuse、Graphiti 和可观测性服务）：
>
> ```bash
> docker compose -f docker-compose.yml -f docker-compose-langfuse.yml -f docker-compose-graphiti.yml -f docker-compose-observability.yml up -d
> ```
>
> 也可以在 shell 中为这些命令设置别名，方便以后快速运行：
>
> ```bash
> alias pentagi="docker compose -f docker-compose.yml -f docker-compose-langfuse.yml -f docker-compose-graphiti.yml -f docker-compose-observability.yml"
> alias pentagi-up="docker compose -f docker-compose.yml -f docker-compose-langfuse.yml -f docker-compose-graphiti.yml -f docker-compose-observability.yml up -d"
> alias pentagi-down="docker compose -f docker-compose.yml -f docker-compose-langfuse.yml -f docker-compose-graphiti.yml -f docker-compose-observability.yml down"
> ```

### 集成知识图谱（Graphiti）

> [!IMPORTANT]
> Graphiti 集成目前仍是**测试版**，对提供商有明显限制。在生产环境中启用前，请先阅读下文的[当前限制](#当前限制)。

PentAGI 集成了由 Neo4j 驱动的时序知识图谱系统 [Graphiti](https://github.com/vxcontrol/pentagi-graphiti)，可理解 AI 智能体运行过程中的语义并跟踪实体关系。vxcontrol 分支增加了渗透测试专用的实体类型和边类型。

#### Graphiti 是什么？

Graphiti 会自动从智能体交互中提取并存储结构化知识，构建包含实体、关系和时间上下文的图谱。借助这些数据，可以实现以下功能：

- **语义记忆**：存储并检索工具、目标、漏洞和技术之间的关系
- **上下文理解**：跟踪不同渗透测试操作随时间形成的关联
- **知识复用**：利用以往渗透测试的经验辅助新的评估
- **高级查询**：搜索复杂模式，例如“哪些工具对相似目标有效？”

#### 启用 Graphiti

Graphiti 知识图谱是**可选功能**，默认关闭。启用步骤如下：

1. 在 `.env` 文件中配置 Graphiti 环境变量：

```bash
## Graphiti 知识图谱设置
GRAPHITI_ENABLED=true
GRAPHITI_TIMEOUT=30
GRAPHITI_URL=http://graphiti:8000
GRAPHITI_MODEL_NAME=gpt-5-mini

# Neo4j 设置（供 Graphiti 服务栈使用）
NEO4J_USER=neo4j
NEO4J_DATABASE=neo4j
NEO4J_PASSWORD=devpassword
NEO4J_URI=bolt://neo4j:7687

# OpenAI API 密钥（Graphiti 提取实体时需要）
OPEN_AI_KEY=your_openai_api_key
```

2. 与 PentAGI 主服务一起启动 Graphiti 服务栈：

```bash
# 如有需要，下载 Graphiti Compose 文件
curl -O https://raw.githubusercontent.com/vxcontrol/pentagi/master/docker-compose-graphiti.yml

# 启动 PentAGI 和 Graphiti
docker compose -f docker-compose.yml -f docker-compose-graphiti.yml up -d
```

3. 验证 Graphiti 是否正常运行：

```bash
# 检查服务健康状态
docker compose -f docker-compose.yml -f docker-compose-graphiti.yml ps graphiti neo4j

# 查看 Graphiti 日志
docker compose -f docker-compose.yml -f docker-compose-graphiti.yml logs -f graphiti

# 访问 Neo4j Browser（可选）
# 打开 http://localhost:7474，使用 NEO4J_USER/NEO4J_PASSWORD 登录

# 访问 Graphiti API（可选，用于调试）
# 打开 http://localhost:8000/docs 查看 Swagger API 文档
```

> [!NOTE]
> Graphiti 服务在 `docker-compose-graphiti.yml` 中定义为独立服务栈。必须同时运行两个 Compose 文件才能启用知识图谱。默认使用预构建 Docker 镜像 `vxcontrol/graphiti:latest`。

#### 存储的内容

启用后，PentAGI 会自动记录以下内容：

- **智能体响应**：智能体的所有推理、分析和决策
- **工具执行记录**：执行的命令、使用的工具及其结果
- **上下文信息**：任务流、任务和子任务的层级关系

#### 当前限制

Graphiti 集成目前仍是测试版。在生产环境中启用前，需要考虑以下限制：

- **仅支持 OpenAI 兼容 LLM。** 随附的 `vxcontrol/graphiti` 镜像只会向一个 OpenAI 兼容端点进行认证。该端点通过 PentAGI `.env` 中的 `OPEN_AI_KEY` 和 `OPEN_AI_SERVER_URL` 配置，后者默认为 `https://api.openai.com/v1`。`docker-compose-graphiti.yml` 会将这两个变量映射为容器内的 `OPENAI_API_KEY` 和 `OPENAI_BASE_URL`，因此无需直接设置容器变量。PentAGI 中为 Anthropic、Google AI（Gemini）、AWS Bedrock、DeepSeek、GLM、Kimi 或 Qwen 配置的凭证**不会**供 Graphiti 提取实体使用。如果部署环境连接不到 OpenAI 兼容端点，请保持 `GRAPHITI_ENABLED=false`。
- **每次部署只能使用一个固定模型。** Graphiti 使用同一个模型名称（`GRAPHITI_MODEL_NAME`，默认为 `gpt-5-mini`）完成所有提取任务，无法按智能体或任务流选择模型。
- **独立计费。** 即使任务流使用的不是 OpenAI 提供商，Graphiti 仍会在所配置的 OpenAI 兼容端点产生费用。
- **暂不提供应用内图谱浏览器。** 如需浏览已记录的图谱，请使用 `http://localhost:7474` 上的 Neo4j Browser 或 `http://localhost:8000/docs` 上的 Graphiti Swagger UI。目前 PentAGI 界面中没有图谱视图。

设置 `GRAPHITI_ENABLED=false` 后，PentAGI 仍会使用主要记忆系统和向量存储正常运行，只跳过额外的知识图谱功能。

### 集成 GitHub 与 Google OAuth

接入 GitHub 和 Google OAuth 后，用户可以使用已有的平台账号完成认证，具体优点包括：

- 无需另建凭证，登录步骤更简单
- 通过可信身份提供商提高安全性
- 可以读取 GitHub/Google 账号的用户资料
- 便于接入现有开发任务流

PentAGI 使用 `PUBLIC_URL` 作为 OAuth 重定向的公开来源和基础 URL。默认部署中，GitHub 与 Google 都通过以下端点处理回调：

```text
${PUBLIC_URL}/api/v1/auth/login-callback
```

GitHub OAuth 配置步骤：

1. 在 GitHub 账号中创建新的 OAuth App。
2. 将 **Homepage URL（主页 URL）** 设为 `PUBLIC_URL` 的值。
3. 将 **Authorization callback URL（授权回调 URL）** 设为 `${PUBLIC_URL}/api/v1/auth/login-callback`。
4. 将客户端凭证写入 `.env` 文件：

```bash
PUBLIC_URL=https://pentagi.example.com
OAUTH_GITHUB_CLIENT_ID=your_github_client_id
OAUTH_GITHUB_CLIENT_SECRET=your_github_client_secret
```

Google OAuth 配置步骤：

1. 在 Google Cloud 项目中创建 OAuth 凭证。
2. 使用相同的回调端点：`${PUBLIC_URL}/api/v1/auth/login-callback`。
3. 将客户端凭证写入 `.env` 文件：

```bash
PUBLIC_URL=https://pentagi.example.com
OAUTH_GOOGLE_CLIENT_ID=your_google_client_id
OAUTH_GOOGLE_CLIENT_SECRET=your_google_client_secret
```

确保 `PUBLIC_URL` 与 PentAGI 实例可从外部访问的 HTTPS 地址一致，并且不要包含回调路径。如果 OAuth 提供商中配置的 URL 与 PentAGI 生成的回调地址不完全一致，提供商会因重定向 URI 不匹配而拒绝登录。

### Docker 镜像配置

PentAGI 可以为不同任务配置 Docker 镜像选择规则。系统会根据任务类型自动选择合适的镜像，也可以通过以下变量限制可选镜像：

| 变量                               | 默认值                          | 说明                                             |
| ---------------------------------- | ------------------------------- | ------------------------------------------------ |
| `PENTAGI_IMAGE`                    | `ghcr.io/fulaoaz/pentagi:zh-cn` | PentAGI 主应用服务使用的汉化镜像                 |
| `DOCKER_DEFAULT_IMAGE`             | `debian:latest`                 | 一般任务或类型不明确的任务所用的默认 Docker 镜像 |
| `DOCKER_DEFAULT_IMAGE_FOR_PENTEST` | `vxcontrol/kali-linux`          | 安全/渗透测试任务所用的默认 Docker 镜像          |

`PENTAGI_IMAGE` 会更改 `docker-compose.yml` 中主 `pentagi` 服务使用的镜像。`DOCKER_DEFAULT_IMAGE` 和 `DOCKER_DEFAULT_IMAGE_FOR_PENTEST` 只影响 PentAGI 内部执行任务时自动选择的工作容器镜像，不会改写 Compose 服务栈的其他部分。因此，`pgvector`、`scraper` 以及可选的 `graphiti` 服务栈仍使用 Compose 文件中定义的镜像。

设置 `DOCKER_DEFAULT_IMAGE` 和 `DOCKER_DEFAULT_IMAGE_FOR_PENTEST` 后，AI 智能体只能从指定镜像中选择。这适用于以下情况：

- **落实安全要求**：只允许使用经过验证且可信的镜像
- **统一环境**：所有操作均使用企业镜像或自定义镜像
- **优化性能**：使用已安装所需工具的预构建镜像

配置示例：

```bash
# 使用自定义 PentAGI 应用镜像
PENTAGI_IMAGE=registry.example.com/security/pentagi:latest

# 为一般任务使用自定义镜像
DOCKER_DEFAULT_IMAGE=mycompany/custom-debian:latest

# 为渗透测试使用专用镜像
DOCKER_DEFAULT_IMAGE_FOR_PENTEST=mycompany/pentest-tools:v2.0
```

> [!NOTE]
> 如果用户在任务中明确指定 Docker 镜像，系统会尝试直接使用该镜像，而忽略这些设置。这些变量只影响系统的自动镜像选择过程。

有关使用自定义渗透测试镜像开展 OpenVAS/GVM 高级实验的方法，请参阅[通过自定义渗透测试镜像使用 OpenVAS](examples/guides/openvas-custom-image.md)。

#### 受限网络、Docker 镜像站与代理

如果当前环境无法直连 Docker Hub（`docker.io`），仅修改 PentAGI 环境变量通常解决不了镜像下载失败。Compose 管理的服务仍依赖 Docker 自身访问镜像仓库，安装器的网络检查也会验证 Docker Hub 是否可达。

在受限网络中，请按以下步骤配置：

1. 确认宿主机可以解析并访问 `docker.io`。
2. 如果 PentAGI 或安装器的 HTTP 流量需要通过出站代理，请设置环境变量 `PROXY_URL`。如需通过代理拉取 Docker 镜像，还要单独配置 Docker 守护进程或 Docker Desktop 代理。Docker 访问镜像仓库时不会使用 PentAGI 的 `PROXY_URL`。
3. 如果 Docker Hub 被阻断或受到严格限流，请先配置组织批准的镜像站或镜像仓库代理，再运行安装器或 `docker compose up`。
4. 更改守护进程配置后重启 Docker，然后重新运行安装器检查或启动 Compose。

Docker 守护进程的镜像站配置示例：

```json
{
  "registry-mirrors": ["https://mirror.example.com"]
}
```

Linux 通常在 `/etc/docker/daemon.json` 中配置此项；Docker Desktop 则使用对应的 Docker Engine 或代理设置。Docker Hub 镜像站可以代理 `vxcontrol/*` 等托管在 Docker Hub 上的镜像，但主 Compose 服务栈还包含 `quay.io/prometheuscommunity/postgres-exporter`，可选的可观测性服务栈也包含 `gcr.io/cadvisor/cadvisor`。这些镜像仓库仍需能够直连，或分别配置获准使用的代理/镜像路径。

详情请参阅 Docker 官方文档中的[镜像仓库镜像站](https://docs.docker.com/docker-hub/image-library/mirror/)和[守护进程代理配置](https://docs.docker.com/engine/daemon/proxy/)。

## 开发

### 开发环境要求

- Go
- Node.js
- Docker
- PostgreSQL
- commitlint

### 环境设置

#### 后端设置

执行一次 `cd backend && go mod download`，安装所需软件包。

运行以下命令生成 Swagger 文件：

```bash
swag init -g ../../pkg/server/router.go -o pkg/server/docs/ --parseDependency --parseInternal --parseDepth 2 -d cmd/pentagi
```

`swag` 的安装命令如下：

```bash
go install github.com/swaggo/swag/cmd/swag@v1.8.7
```

运行以下命令生成 GraphQL resolver 文件：

```bash
go run github.com/99designs/gqlgen --config ./gqlgen/gqlgen.yml
```

生成的文件位于 `pkg/graph` 目录。

根据 sqlc 配置生成 ORM 方法（数据库包）：

```bash
docker run --rm -v $(pwd):/src -w /src --network pentagi-network -e DATABASE_URL="{URL}" sqlc/sqlc:1.27.0 generate -f sqlc/sqlc.yml
```

根据 OpenAPI 规范生成 Langfuse SDK：

```bash
fern generate --local
```

安装 fern-cli：

```bash
pnpm add -g fern-api
```

#### 测试

运行测试：`cd backend && go test -v ./...`

#### 前端设置

执行一次 `cd frontend && pnpm install`，安装所需软件包。

运行 `pnpm run graphql:generate`，根据 `graphql-codegen.ts` 生成 GraphQL 文件。

确保已全局安装 `graphql-codegen`：

```bash
pnpm add -g graphql-codegen
```

然后可以运行：
* `pnpm run prettier`：检查代码格式
* `pnpm run prettier:fix`：修复代码格式
* `pnpm run lint`：检查代码是否通过 lint
* `pnpm run lint:fix`：修复 lint 问题

运行 `pnpm run ssl:generate` 可通过 `generate-ssl.ts` 生成 SSL 证书；如果直接运行 `pnpm run dev`，系统会自动生成证书。

#### 后端配置

在 `.vscode/launch.json` 文件中修改 `backend` 配置：
- `DATABASE_URL`：PostgreSQL 数据库 URL（例如 `postgres://postgres:postgres@localhost:5432/pentagidb?sslmode=disable`）
- `DOCKER_HOST`：Docker SDK API（macOS 示例：`DOCKER_HOST=unix:///Users/<my-user>/Library/Containers/com.docker.docker/Data/docker.raw.sock`），参见[更多信息](https://stackoverflow.com/a/62757128/5922857)

可选配置：
- `SERVER_PORT`：服务器运行端口（默认：`8443`）
- `SERVER_USE_SSL`：是否为服务器启用 SSL（默认：`false`）

##### PostgreSQL/pgvector 连接池大小

PentAGI 会为同一个 PostgreSQL 实例打开两个相互独立的连接池：

| 连接池 | 环境变量 | 默认值 | 用途 |
|---|---|---|---|
| 共享 `sql.DB` | `DATABASE_MAX_OPEN_CONNS` | `25` | 所有 sqlc 查询和 GORM 处理程序共用一个 `*sql.DB` |
| 共享 `pgxpool` | `DATABASE_VECTOR_MAX_CONNS` | `10` | 所有 pgvector 存储（智能体记忆 + 知识 API）共用一个连接池 |

其他调优参数：
- `DATABASE_MAX_IDLE_CONNS`：两次请求之间，`sql.DB` 池中保持打开状态的最大空闲连接数（默认：`5`）。

**标准 `vxcontrol/pgvector` 镜像的连接预算**（`max_connections = 100`，`superuser_reserved_connections = 3`）：

```
客户端可用连接数                        = 97
  pentagi sql.DB  (DATABASE_MAX_OPEN_CONNS)   = 25
  pentagi pgxpool (DATABASE_VECTOR_MAX_CONNS) = 10
  pgexporter                                  =  3
  autovacuum 工作进程                         =  3
  ─────────────────────────────────────────
  已使用连接总数                              = 41
  空闲余量                                    = 56  (≈ 58 %)
```

默认值按 **10 个并行任务流**及并发 API 请求进行配置。如果需要运行更多任务流，或让多个 PentAGI 实例共用同一个 PostgreSQL，请在 `docker-compose.yml` 中通过 `command` 覆盖项提高 `max_connections`，并按比例增大连接池：

```yaml
pgvector:
  image: vxcontrol/pgvector:latest
  command: postgres -c max_connections=200
```

检查正在运行的部署当前使用的连接预算：

```bash
# PostgreSQL 限制
docker exec pgvector sh -c 'psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c \
  "SELECT name, setting FROM pg_settings
   WHERE name IN ('"'"'max_connections'"'"', '"'"'superuser_reserved_connections'"'"');"'

# 当前用量与可用量
docker exec pgvector sh -c 'psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c \
  "SELECT max_conn, used, max_conn - used AS available
   FROM (SELECT current_setting('"'"'max_connections'"'"')::int AS max_conn,
                count(*) AS used FROM pg_stat_activity) t;"'

# 按客户端细分
docker exec pgvector sh -c 'psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c \
  "SELECT application_name, client_addr, state, count(*)
   FROM pg_stat_activity
   WHERE pid <> pg_backend_pid()
   GROUP BY 1, 2, 3 ORDER BY count DESC;"'
```

#### 前端配置

在 `.vscode/launch.json` 文件中修改 `frontend` 配置：
- `VITE_API_URL`：后端 API URL。*不要包含* URL 协议部分（例如应填 `localhost:8080`，*不要填* `http://localhost:8080`）
- `VITE_USE_HTTPS`：是否为服务器启用 SSL（默认：`false`）
- `VITE_PORT`：服务器运行端口（默认：`8000`）
- `VITE_HOST`：服务器监听地址（默认：`0.0.0.0`）

### 运行应用

#### 后端

在 `backend` 目录中运行以下命令：
- 使用 `.env` 文件设置环境变量，例如执行 `source .env`
- 执行 `go run cmd/pentagi/main.go` 启动服务器

> [!NOTE]
> 首次运行需要下载依赖项和 Docker 镜像来设置后端环境，因此可能耗时较长。

#### 前端

在 `frontend` 目录中运行以下命令：
- 执行 `pnpm install` 安装依赖项
- 执行 `pnpm run dev` 运行网页应用
- 执行 `pnpm run build` 构建网页应用

在浏览器中打开网页应用的 URL。

## 测试 LLM 智能体

PentAGI 提供 `ctester` 工具，用于测试和验证 LLM 智能体的能力。它可以检查 LLM 提供商配置能否适配不同类型的智能体，并帮助为各个角色选择合适的模型。

该工具支持并行测试多个智能体、生成详细报告，并提供灵活的配置选项。

### 主要功能

- **并行测试**：同时测试多个智能体，更快获得结果
- **完整的测试套件**：评估基本补全、JSON 响应、函数调用和渗透测试知识
- **详细报告**：生成包含成功率和性能指标的 Markdown 报告
- **灵活配置**：按需选择特定智能体或测试组
- **专项测试组**：包含针对网络安全和渗透测试场景的领域测试

### 使用场景

#### 开发人员（使用本地 Go 环境）

如果已经克隆仓库并安装 Go，可以使用以下命令：

```bash
# 使用 .env 文件中的默认配置
cd backend
go run cmd/ctester/*.go -verbose

# 自定义提供商配置
go run cmd/ctester/*.go -config ../examples/configs/openrouter.provider.yml -verbose

# 生成报告文件
go run cmd/ctester/*.go -config ../examples/configs/deepinfra.provider.yml -report ../test-report.md

# 仅测试指定类型的智能体
go run cmd/ctester/*.go -agents simple,simple_json,primary_agent -verbose

# 仅运行指定测试组
go run cmd/ctester/*.go -groups basic,advanced -verbose
```

#### 普通用户（使用 Docker 镜像）

如果不想搭建开发环境，可以直接使用预构建 Docker 镜像：

```bash
# 使用 Docker 和默认环境进行测试
docker run --rm -v $(pwd)/.env:/opt/pentagi/.env vxcontrol/pentagi /opt/pentagi/bin/ctester -verbose

# 使用自定义提供商配置进行测试
docker run --rm \
  -v $(pwd)/.env:/opt/pentagi/.env \
  -v $(pwd)/my-config.yml:/opt/pentagi/config.yml \
  vxcontrol/pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/config.yml -agents simple,primary_agent,coder -verbose

# 生成详细报告
docker run --rm \
  -v $(pwd)/.env:/opt/pentagi/.env \
  -v $(pwd):/opt/pentagi/output \
  vxcontrol/pentagi /opt/pentagi/bin/ctester -report /opt/pentagi/output/report.md
```

#### 使用预配置提供商

Docker 镜像内置了对主要提供商（OpenAI、Anthropic、Gemini、Ollama）的支持，还包含其他服务（OpenRouter、DeepInfra、DeepSeek、Moonshot、Novita）的预配置文件：

```bash
# 测试 OpenRouter 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/openrouter.provider.yml

# 测试 DeepInfra 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/deepinfra.provider.yml

# 测试 DeepSeek 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -provider deepseek

# 测试 GLM 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -provider glm

# 测试 Kimi 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -provider kimi

# 测试 Qwen 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -provider qwen

# 使用自定义提供商的 DeepSeek 配置文件进行测试
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/deepseek.provider.yml

# 使用自定义提供商的 Moonshot 配置文件进行测试
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/moonshot.provider.yml

# 测试 Novita 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/novita.provider.yml

# 测试 OpenAI 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -type openai

# 测试 Anthropic 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -type anthropic

# 测试 Gemini 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -type gemini

# 测试 AWS Bedrock 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -type bedrock

# 测试自定义 OpenAI 配置
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/custom-openai.provider.yml

# 测试 Ollama 配置（本地推理）
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/ollama-llama318b.provider.yml

# 测试 Ollama Qwen3 32B 配置（需要创建自定义模型）
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/ollama-qwen332b-fp16-tc.provider.yml

# 测试 Ollama QwQ 32B 配置（需要创建自定义模型，并准备 71.3 GB 显存）
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/ollama-qwq32b-fp16-tc.provider.yml
```

要使用这些配置，`.env` 文件只需包含以下内容：

```
LLM_SERVER_URL=https://openrouter.ai/api/v1      # 或 https://api.deepinfra.com/v1/openai、https://api.openai.com/v1、https://api.novita.ai/openai
LLM_SERVER_KEY=your_api_key
LLM_SERVER_MODEL=                                # 留空，模型已在配置文件中指定
LLM_SERVER_CONFIG_PATH=/opt/pentagi/conf/openrouter.provider.yml  # 或 deepinfra.provider.ymll、custom-openai.provider.yml、novita.provider.yml
LLM_SERVER_PROVIDER=                             # LiteLLM 代理使用的提供商名称（例如 openrouter、deepseek、moonshot、novita）
LLM_SERVER_LEGACY_REASONING=false                # 控制推理格式；OpenAI 必须设为 true（默认：false）
LLM_SERVER_PRESERVE_REASONING=false              # 在多轮对话中保留推理内容（Moonshot 要求开启，默认：false）

# OpenAI（官方 API）
OPEN_AI_KEY=your_openai_api_key                  # OpenAI API 密钥
OPEN_AI_SERVER_URL=https://api.openai.com/v1     # OpenAI API 端点

# Anthropic（Claude 模型）
ANTHROPIC_API_KEY=your_anthropic_api_key         # Anthropic API 密钥
ANTHROPIC_SERVER_URL=https://api.anthropic.com/v1  # Anthropic API 端点

# Gemini（Google AI）
GEMINI_API_KEY=your_gemini_api_key               # Google AI API 密钥
GEMINI_SERVER_URL=https://generativelanguage.googleapis.com  # Google AI API 端点

# AWS Bedrock（企业基础模型）
BEDROCK_REGION=us-east-1                         # Bedrock 服务所在的 AWS 区域
# 认证（选择一种方式，优先级：DefaultAuth > BearerToken > AccessKey）：
BEDROCK_DEFAULT_AUTH=false                       # 使用 AWS SDK 凭证链（环境变量、EC2 角色、~/.aws/credentials）
BEDROCK_BEARER_TOKEN=                            # Bearer 令牌认证（优先级高于静态凭证）
BEDROCK_ACCESS_KEY_ID=your_aws_access_key        # AWS 访问密钥 ID（静态凭证）
BEDROCK_SECRET_ACCESS_KEY=your_aws_secret_key    # AWS 秘密访问密钥（静态凭证）
BEDROCK_SESSION_TOKEN=                           # AWS 会话令牌（可选，与静态认证配合使用的临时凭证）
BEDROCK_SERVER_URL=                              # 可选的自定义 Bedrock 端点（VPC 端点、本地测试）

# Ollama（本地服务器或云服务）
OLLAMA_SERVER_URL=                               # 本地：http://ollama-server:11434，云服务：https://ollama.com
OLLAMA_SERVER_API_KEY=                           # Ollama Cloud 需要填写（https://ollama.com/settings/keys），本地使用时留空
OLLAMA_SERVER_MODEL=
OLLAMA_SERVER_CONFIG_PATH=
OLLAMA_SERVER_PULL_MODELS_TIMEOUT=
OLLAMA_SERVER_PULL_MODELS_ENABLED=
OLLAMA_SERVER_LOAD_MODELS_ENABLED=

# DeepSeek（推理能力较强的中文 AI）
DEEPSEEK_API_KEY=                                # DeepSeek API 密钥
DEEPSEEK_SERVER_URL=https://api.deepseek.com     # DeepSeek API 端点
DEEPSEEK_PROVIDER=                               # 可选：LiteLLM 前缀（例如 'deepseek'）

# GLM（智谱 AI）
GLM_API_KEY=                                     # GLM API 密钥
GLM_SERVER_URL=https://api.z.ai/api/paas/v4      # GLM API 端点（国际站）
GLM_PROVIDER=                                    # 可选：LiteLLM 前缀（例如 'zai'）

# Kimi（Moonshot AI）
KIMI_API_KEY=                                    # Kimi API 密钥
KIMI_SERVER_URL=https://api.moonshot.ai/v1       # Kimi API 端点（国际站）
KIMI_PROVIDER=                                   # 可选：LiteLLM 前缀（例如 'moonshot'）

# Qwen（阿里云 DashScope）
QWEN_API_KEY=                                    # Qwen API 密钥
QWEN_SERVER_URL=https://dashscope-us.aliyuncs.com/compatible-mode/v1  # Qwen API 端点（美国）
QWEN_PROVIDER=                                   # 可选：LiteLLM 前缀（例如 'dashscope'）

# Ollama（本地推理）使用上方变量
OLLAMA_SERVER_URL=http://localhost:11434
OLLAMA_SERVER_MODEL=llama3.1:8b-instruct-q8_0
OLLAMA_SERVER_CONFIG_PATH=/opt/pentagi/conf/ollama-llama318b.provider.yml
OLLAMA_SERVER_PULL_MODELS_ENABLED=false
OLLAMA_SERVER_LOAD_MODELS_ENABLED=false
```

#### 在组织未验证时使用 OpenAI

如果 OpenAI 账号所属组织尚未验证，且尚未获得最新推理模型（o1、o3、o4-mini）的访问权限，需要改用自定义配置。

使用此类 OpenAI 账号时，请按以下内容配置 `.env` 文件：

```bash
LLM_SERVER_URL=https://api.openai.com/v1
LLM_SERVER_KEY=your_openai_api_key
LLM_SERVER_MODEL=                                # 留空，模型已在配置文件中指定
LLM_SERVER_CONFIG_PATH=/opt/pentagi/conf/custom-openai.provider.yml
LLM_SERVER_LEGACY_REASONING=true                 # OpenAI 推理格式要求开启
```

此配置使用预构建的 `custom-openai.provider.yml` 文件，将所有智能体类型映射到组织未验证时可用的模型，并用 `o3-mini` 替代 `o1`、`o3` 和 `o4-mini` 等模型。

可以运行以下命令测试配置：

```bash
# 使用适合未验证账号的自定义 OpenAI 配置进行测试
docker exec -it pentagi /opt/pentagi/bin/ctester -config /opt/pentagi/conf/custom-openai.provider.yml
```

> [!NOTE]
> 必须设置 `LLM_SERVER_LEGACY_REASONING=true` 才能兼容 OpenAI，它会按 OpenAI API 要求的格式发送推理参数。

#### 使用 LiteLLM 代理

通过 LiteLLM 代理访问不同 LLM 提供商时，模型名称前会加上提供商名称（例如使用 `moonshot/kimi-2.5`，而不是 `kimi-2.5`）。如果希望同一份提供商配置文件既能直连 API，也能通过 LiteLLM 代理使用，请设置 `LLM_SERVER_PROVIDER` 变量：

```bash
# 直连 Moonshot API
LLM_SERVER_URL=https://api.moonshot.ai/v1
LLM_SERVER_KEY=your_moonshot_api_key
LLM_SERVER_CONFIG_PATH=/opt/pentagi/conf/moonshot.provider.yml
LLM_SERVER_PROVIDER=                             # 直连时留空

# 通过 LiteLLM 代理访问
LLM_SERVER_URL=http://litellm-proxy:4000
LLM_SERVER_KEY=your_litellm_api_key
LLM_SERVER_CONFIG_PATH=/opt/pentagi/conf/moonshot.provider.yml
LLM_SERVER_PROVIDER=moonshot                     # LiteLLM 的提供商前缀
```

设置 `LLM_SERVER_PROVIDER=moonshot` 后，系统会自动为配置文件中的所有模型名称添加 `moonshot/` 前缀，使其符合 LiteLLM 的模型命名规则。

**LiteLLM 提供商名称映射：**

使用 LiteLLM 代理时，设置对应的 `*_PROVIDER` 变量即可启用模型前缀：

- `deepseek`：用于 DeepSeek 模型（`DEEPSEEK_PROVIDER=deepseek` → `deepseek/deepseek-v4-flash`）
- `zai`：用于 GLM 模型（`GLM_PROVIDER=zai` → `zai/glm-4`）
- `moonshot`：用于 Kimi 模型（`KIMI_PROVIDER=moonshot` → `moonshot/kimi-k2.5`）
- `dashscope`：用于 Qwen 模型（`QWEN_PROVIDER=dashscope` → `dashscope/qwen-plus`）
- `openai`、`anthropic`、`gemini`：用于主要云服务提供商
- `openrouter`：用于 OpenRouter 聚合服务
- `deepinfra`：用于 DeepInfra 托管服务
- `novita`：用于 Novita AI
- LiteLLM 实例中配置的其他任意提供商名称

**LiteLLM 示例：**
```bash
# 通过 LiteLLM 代理使用 DeepSeek 模型，并添加模型前缀
DEEPSEEK_API_KEY=your_litellm_proxy_key
DEEPSEEK_SERVER_URL=http://litellm-proxy:4000
DEEPSEEK_PROVIDER=deepseek  # 在 LiteLLM 中，模型名称变为 deepseek/deepseek-v4-flash、deepseek/deepseek-v4-pro

# 直连 DeepSeek API（无需前缀）
DEEPSEEK_API_KEY=your_deepseek_api_key
DEEPSEEK_SERVER_URL=https://api.deepseek.com
# DEEPSEEK_PROVIDER 留空
```

这种方式可以：
- 让直连和代理访问共用同一份配置文件
- 无需修改配置文件即可切换提供商
- 通过 LiteLLM 方便地测试不同路由策略

#### 在生产环境中运行测试

如果已有正在运行的 PentAGI 容器，可以按以下方式测试当前配置：

```bash
# 在现有容器中使用当前环境变量运行 ctester
docker exec -it pentagi /opt/pentagi/bin/ctester -verbose

# 按固定顺序测试指定类型的智能体
docker exec -it pentagi /opt/pentagi/bin/ctester -agents simple,primary_agent,pentester -groups basic,knowledge -verbose

# 在容器内生成报告文件
docker exec -it pentagi /opt/pentagi/bin/ctester -report /opt/pentagi/data/agent-test-report.md

# 从宿主机获取报告
docker cp pentagi:/opt/pentagi/data/agent-test-report.md ./
```

### 命令行选项

该工具支持以下选项：

- `-env <path>`：环境文件路径（默认：`.env`）
- `-type <provider>`：提供商类型，可选 `custom`、`openai`、`anthropic`、`ollama`、`bedrock`、`gemini`（默认：`custom`）
- `-config <path>`：自定义提供商配置的路径（默认读取环境变量 `LLM_SERVER_CONFIG_PATH`）
- `-tests <path>`：自定义测试 YAML 文件的路径（可选）
- `-report <path>`：报告文件的写入路径（可选）
- `-agents <list>`：要测试的智能体类型，以逗号分隔（默认：`all`）
- `-groups <list>`：要运行的测试组，以逗号分隔（默认：`all`）
- `-verbose`：启用详细输出，显示各智能体的具体测试结果

### 可用的智能体类型

智能体按以下固定顺序进行测试：

1. **simple**：基本补全任务
2. **simple_json**：JSON 结构化响应
3. **primary_agent**：主推理智能体
4. **assistant**：交互式助手模式
5. **generator**：内容生成
6. **refiner**：内容优化和改进
7. **adviser**：专家建议和咨询
8. **reflector**：自我反思和分析
9. **searcher**：信息收集和搜索
10. **enricher**：数据补充和扩展
11. **coder**：代码生成和分析
12. **installer**：安装和设置任务
13. **pentester**：渗透测试和安全评估

### 可用的测试组

- **basic**：基本补全和提示词响应测试
- **advanced**：复杂推理和函数调用测试
- **json**：JSON 格式验证和结构测试（专为 `simple_json` 智能体设计）
- **knowledge**：网络安全和渗透测试领域的专项知识测试

> **注意**：`json` 测试组专为 `simple_json` 智能体设计，其他所有智能体使用 `basic`、`advanced` 和 `knowledge` 测试组。这样的划分可以覆盖各智能体的预期用途。

### 提供商配置示例

提供商配置用于指定不同类型智能体所使用的模型：

```yaml
simple:
  model: "provider/model-name"
  temperature: 0.7
  top_p: 0.95
  n: 1
  max_tokens: 4000

simple_json:
  model: "provider/model-name"
  temperature: 0.7
  top_p: 1.0
  n: 1
  max_tokens: 4000
  json: true

# ... 其他智能体类型 ...
```

### 优化流程

1. **建立基线**：使用默认配置运行测试，记录基准性能
2. **分析各智能体的性能**：按固定测试顺序检查结果，找出表现较差的智能体
3. **测试专项配置**：通过提供商专用配置，为各类智能体尝试不同模型
4. **关注领域知识**：重点检查用于评估网络安全知识的 `knowledge` 测试组
5. **验证函数调用**：确认关键智能体类型能够稳定通过工具类测试
6. **比较结果**：比较所有测试组的成功率和性能，选出效果最好的方案
7. **部署最优配置**：将优化后的设置用于生产环境

借助该工具，可以为不同 AI 智能体选择更适合其任务的模型，在提高可靠性的同时控制成本。

## 嵌入模型配置与测试

PentAGI 使用向量嵌入实现语义搜索、知识存储和记忆管理。系统支持多种嵌入模型提供商，可按实际需求进行配置。

### 支持的嵌入模型提供商

PentAGI 支持以下嵌入模型提供商：

- **OpenAI**（默认）：使用 OpenAI 文本嵌入模型
- **Ollama**：通过 Ollama 运行本地嵌入模型
- **Mistral**：使用 Mistral AI 嵌入模型
- **Jina**：使用 Jina AI 嵌入服务
- **HuggingFace**：使用 HuggingFace 模型
- **GoogleAI**：使用 Google 嵌入模型
- **VoyageAI**：使用 VoyageAI 嵌入模型

> **兼容 OpenAI 的第三方服务**：任何提供 OpenAI `/embeddings` API 的服务都可以接入。设置 `EMBEDDING_PROVIDER=openai`，并在 `EMBEDDING_URL` 中填写自定义地址即可。例如，**Qwen DashScope** 通过 `/compatible-mode/v1` 端点提供 `text-embedding-v4`，但仅限国际站和中国大陆区域，美国区域没有嵌入接口。完整配置片段参见 Qwen 章节的[其他集成方式](#其他集成方式)。

<details>
<summary><b>嵌入模型提供商配置</b>（点击展开）</summary>

### 环境变量

在 `.env` 文件中设置以下环境变量，配置嵌入模型提供商：

```bash
# 主要嵌入配置
EMBEDDING_PROVIDER=openai       # 提供商类型（openai、ollama、mistral、jina、huggingface、googleai、voyageai）
EMBEDDING_MODEL=text-embedding-3-small  # 使用的模型名称
EMBEDDING_URL=                  # 可选的自定义 API 端点
EMBEDDING_KEY=                  # 提供商 API 密钥（如需要）
EMBEDDING_BATCH_SIZE=100        # 每批处理的文档数
EMBEDDING_STRIP_NEW_LINES=true  # 嵌入前是否移除文本中的换行符
EMBEDDING_MAX_TEXT_BYTES=8192   # 每份文档发送给嵌入模型的最大字节数（以字节数近似限制令牌量）

# 高级设置
PROXY_URL=                      # 所有 API 调用使用的可选代理
HTTP_CLIENT_TIMEOUT=600         # 外部 API 调用超时时间，单位为秒（默认：600；0 表示不超时）
TERMINAL_TOOL_TIMEOUT=1200      # 终端工具命令在 timeout=0 或负数时使用的默认超时秒数（范围：1～10800；小于等于 0 或超过 10800 的值会限制为 10800，即 3 小时）

# SSL/TLS 证书配置（用于同 LLM 后端和工具服务器进行外部通信）
EXTERNAL_SSL_CA_PATH=           # 容器内自定义 CA 证书文件的路径（PEM 格式）
                                # 必须指向 /opt/pentagi/ssl/ 目录（例如 /opt/pentagi/ssl/ca-bundle.pem）
EXTERNAL_SSL_INSECURE=false     # 跳过证书验证（仅用于测试）
```

<details>
<summary><b>如何添加自定义 CA 证书</b>（点击展开）</summary>

如果遇到以下错误：`tls: failed to verify certificate: x509: certificate signed by unknown authority`

**第 1 步：** 获取 PEM 格式的 CA 证书包（可包含多个证书）。

**第 2 步：** 将文件放入宿主机的 SSL 目录：
```bash
# 默认位置（未设置 PENTAGI_SSL_DIR 时）
cp ca-bundle.pem ./pentagi-ssl/

# 或自定义位置（在 docker-compose.yml 中使用 PENTAGI_SSL_DIR 时）
cp ca-bundle.pem /path/to/your/ssl/dir/
```

**第 3 步：** 在 `.env` 文件中设置路径（必须是容器内路径）：
```bash
# pentagi-ssl 卷挂载到容器内的 /opt/pentagi/ssl
EXTERNAL_SSL_CA_PATH=/opt/pentagi/ssl/ca-bundle.pem
EXTERNAL_SSL_INSECURE=false
```

**第 4 步：** 重启 PentAGI：
```bash
docker compose restart pentagi
```

**注意事项：**
- `pentagi-ssl` 卷挂载到容器内的 `/opt/pentagi/ssl`
- 可以通过 docker-compose.yml 中的 `PENTAGI_SSL_DIR` 变量更改宿主机目录
- 一个 PEM 文件可以包含多个证书和中间 CA
- `EXTERNAL_SSL_INSECURE=true` 仅用于测试，不建议在生产环境中使用

</details>

### 各提供商的限制

各提供商支持的功能和限制如下：

- **OpenAI**：支持所有配置选项
- **Ollama**：使用本地模型，因此不使用 `EMBEDDING_KEY`
- **Mistral**：不支持 `EMBEDDING_MODEL` 或自定义 HTTP 客户端
- **Jina**：不支持自定义 HTTP 客户端
- **HuggingFace**：必须设置 `EMBEDDING_KEY`，支持其他所有选项
- **GoogleAI**：不支持 `EMBEDDING_URL`，必须设置 `EMBEDDING_KEY`
- **VoyageAI**：支持所有配置选项

如果未指定 `EMBEDDING_URL` 和 `EMBEDDING_KEY`，系统会尝试使用对应 LLM 提供商的设置。例如，`EMBEDDING_PROVIDER=openai` 时使用 `OPEN_AI_KEY`。

### 为什么要保持嵌入模型提供商一致

必须持续使用同一个嵌入模型提供商，原因如下：

1. **向量兼容性**：不同提供商生成的向量在维度和数学性质上有所不同
2. **语义一致性**：更换提供商会破坏之前已嵌入文档之间的语义相似关系
3. **记忆数据失效**：混用不同嵌入向量会降低搜索质量，并使知识库功能异常

更换嵌入模型提供商后，应清空整个知识库并重新建立索引（参见下文的 `etester` 工具）。

</details>

### 嵌入测试工具（etester）

PentAGI 提供专用的 `etester` 工具，用于测试、管理和调试嵌入功能，以及排查向量嵌入与知识存储问题。

<details>
<summary><b>etester 命令</b>（点击展开）</summary>

```bash
# 测试嵌入模型提供商和数据库连接
cd backend
go run cmd/etester/main.go test -verbose

# 显示嵌入数据库的统计信息
go run cmd/etester/main.go info

# 删除嵌入数据库中的所有文档（请谨慎操作）
go run cmd/etester/main.go flush

# 为所有文档重新计算嵌入向量（更换提供商后使用）
go run cmd/etester/main.go reindex

# 在嵌入数据库中搜索文档
go run cmd/etester/main.go search -query "如何安装 PostgreSQL" -limit 5
```

### 使用 Docker

如果 PentAGI 在 Docker 中运行，可以从容器内使用 etester：

```bash
# 测试嵌入模型提供商
docker exec -it pentagi /opt/pentagi/bin/etester test

# 显示详细的数据库信息
docker exec -it pentagi /opt/pentagi/bin/etester info -verbose
```

### 高级搜索选项

`search` 命令支持多种筛选条件，可缩小结果范围：

```bash
# 按文档类型筛选
docker exec -it pentagi /opt/pentagi/bin/etester search -query "安全漏洞" -doc_type guide -threshold 0.8

# 按任务流 ID 筛选
docker exec -it pentagi /opt/pentagi/bin/etester search -query "代码示例" -doc_type code -flow_id 42

# 所有可用搜索选项
docker exec -it pentagi /opt/pentagi/bin/etester search -help
```

可用的搜索参数：
- `-query STRING`：搜索查询文本（必填）
- `-doc_type STRING`：按文档类型筛选（answer、memory、guide、code）
- `-flow_id NUMBER`：按任务流 ID 筛选（正数）
- `-answer_type STRING`：按答案类型筛选（guide、vulnerability、code、tool、other）
- `-guide_type STRING`：按指南类型筛选（install、configure、use、pentest、development、other）
- `-limit NUMBER`：最大结果数（默认：3）
- `-threshold NUMBER`：相似度阈值（0.0～1.0，默认：0.7）

### 跨任务流的记忆生命周期

PentAGI 存储多种向量文档，各自用途不同：

- `memory` 记录任务流专属的执行历史，例如工具结果和智能体观察
- `guide`、`answer` 和 `code` 用于存储可复用知识，为后续运行提供帮助

如需查看某次评估的执行过程，请使用对应的 `flow_id` 搜索向量存储。如果希望知识在单次运行结束后仍可复用，应将长期有效的结果明确保存为 `guide`、`answer` 或 `code` 文档，不要只依赖执行记忆。

例如，如果某个目标有需要反复使用的设置说明、认证注意事项或目标专用测试方法，可以要求智能体将这些信息保存为 `guide`，并在下一次评估开始时进行搜索。如果希望新任务流从可复用上下文开始，目前建议采用这种工作方式。

删除任务流时，PentAGI 的软删除机制会使其不再出现在常规查询中，因此可复用知识应与单个任务流的执行历史分开管理。如果启用了本文前述的可选 Graphiti 知识图谱，除非另外构建跨任务流复用机制，否则应将其当前搜索上下文视为仅限正在进行的任务流或评估。

### 常见故障排查场景

1. **更换嵌入模型提供商后**：务必运行 `flush` 或 `reindex`，保持数据一致
2. **搜索结果不理想**：尝试调整相似度阈值，或检查嵌入向量是否正确生成
3. **数据库连接问题**：确认 PostgreSQL 正在运行且已安装 pgvector 扩展
4. **缺少 API 密钥**：检查所选嵌入模型提供商的环境变量

</details>

## 使用 ftester 测试函数

PentAGI 提供 `ftester` 工具，用于调试、测试和开发特定函数及 AI 智能体行为。`ctester` 侧重测试 LLM 模型能力，而 `ftester` 可以直接调用单个系统函数和 AI 智能体组件，并精确控制执行上下文。

### 主要功能

- **直接访问函数**：无需运行整个系统即可测试单个函数
- **模拟模式**：通过内置模拟对象测试函数，无需使用正在运行的 PentAGI 部署
- **交互式输入**：通过交互方式填写函数参数，进行探索性测试
- **详细输出**：使用颜色区分终端输出，并格式化响应和错误
- **上下文感知测试**：在指定任务流、任务和子任务的上下文中调试 AI 智能体
- **可观测性集成**：所有函数调用都会记录到 Langfuse 和可观测性服务栈

### 使用模式

#### 命令行参数

可以直接从命令行运行 ftester，并指定函数和参数：

```bash
# 模拟模式的基本用法
cd backend
go run cmd/ftester/main.go [function_name] -[arg1] [value1] -[arg2] [value2]

# 示例：在模拟模式下测试终端命令
go run cmd/ftester/main.go terminal -command "ls -la" -message "列出文件"

# 使用真实任务流上下文
go run cmd/ftester/main.go -flow 123 terminal -command "whoami" -message "检查用户"

# 在指定任务/子任务上下文中测试 AI 智能体
go run cmd/ftester/main.go -flow 123 -task 456 -subtask 789 pentester -message "查找漏洞"
```

#### 交互模式

不带参数运行 ftester，即可进入引导式交互模式：

```bash
# 启动交互模式
go run cmd/ftester/main.go [function_name]

# 示例：以交互方式填写浏览器工具参数
go run cmd/ftester/main.go browser
```

<details>
<summary><b>可用函数</b>（点击展开）</summary>

### 环境函数
- **terminal**：在容器中执行命令并返回输出
- **file**：在容器中执行文件操作（读取、写入、列出）

### 搜索函数
- **browser**：访问网站并捕获截图
- **google**：使用 Google 自定义搜索来搜索网页
- **duckduckgo**：使用 DuckDuckGo 搜索网页
- **tavily**：使用 Tavily AI 搜索引擎
- **traversaal**：使用 Traversaal AI 搜索引擎
- **perplexity**：使用 Perplexity AI 搜索
- **sploitus**：搜索安全漏洞利用、漏洞（CVE）和渗透测试工具
- **searxng**：使用 Searxng 元搜索引擎，聚合多个搜索引擎的结果

### 向量数据库函数
- **search_in_memory**：在向量数据库中搜索信息
- **search_guide**：在向量数据库中查找指南文档
- **search_answer**：在向量数据库中查找问题答案
- **search_code**：在向量数据库中查找代码示例

### AI 智能体函数
- **advice**：从 AI 智能体获取专家建议
- **coder**：请求生成或修改代码
- **maintenance**：运行系统维护任务
- **memorist**：在向量数据库中存储和整理信息
- **pentester**：执行安全测试和漏洞分析
- **search**：在多个来源中执行复杂搜索

### 实用函数
- **describe**：显示任务流、任务和子任务的信息

</details>

<details>
<summary><b>调试任务流上下文</b>（点击展开）</summary>

`describe` 函数提供任务流内各任务和子任务的详细信息，适合在 PentAGI 出现问题或卡住时进行诊断。

```bash
# 列出系统中的所有任务流
go run cmd/ftester/main.go describe

# 显示指定任务流的所有任务和子任务
go run cmd/ftester/main.go -flow 123 describe

# 显示指定任务的详细信息
go run cmd/ftester/main.go -flow 123 -task 456 describe

# 显示指定子任务的详细信息
go run cmd/ftester/main.go -flow 123 -task 456 -subtask 789 describe

# 显示包含完整说明和结果的详细输出
go run cmd/ftester/main.go -flow 123 describe -verbose
```

该函数可以定位任务流卡住的具体位置，再通过直接调用相应的智能体函数来恢复处理。

</details>

<details>
<summary><b>函数帮助与发现</b>（点击展开）</summary>

每个函数都提供帮助模式，用于显示可用参数：

```bash
# 获取指定函数的帮助
go run cmd/ftester/main.go [function_name] -help

# 示例：
go run cmd/ftester/main.go terminal -help
go run cmd/ftester/main.go browser -help
go run cmd/ftester/main.go describe -help
```

也可以不带参数运行 ftester，查看所有可用函数：

```bash
go run cmd/ftester/main.go
```

</details>

<details>
<summary><b>输出格式</b>（点击展开）</summary>

`ftester` 使用不同颜色标记输出，方便阅读：

- **蓝色标题**：章节标题和键名
- **青色 [INFO]**：一般信息
- **绿色 [SUCCESS]**：操作成功
- **红色 [ERROR]**：错误信息
- **黄色 [WARNING]**：警告信息
- **黄色 [MOCK]**：表示模拟模式操作
- **品红色值**：函数参数和结果

JSON 和 Markdown 响应会自动格式化，便于阅读。

</details>

<details>
<summary><b>高级使用场景</b>（点击展开）</summary>

### 调试卡住的 AI 任务流

PentAGI 在任务流中卡住时：

1. 通过界面暂停任务流
2. 使用 `describe` 确定当前任务和子任务
3. 使用相同的任务/子任务 ID 直接调用智能体函数
4. 检查详细输出，定位问题
5. 根据需要恢复任务流或手动干预

### 测试环境变量

验证 API 密钥和外部服务是否配置正确：

```bash
# 测试 Google 搜索 API 配置
go run cmd/ftester/main.go google -query "渗透测试工具"

# 测试浏览器能否访问外部网站
go run cmd/ftester/main.go browser -url "https://example.com"
```

### 开发新的 AI 智能体行为

开发新的提示词模板或智能体行为时：

1. 在界面中创建测试任务流
2. 使用 ftester 直接调用智能体，并尝试不同提示词
3. 观察响应并相应调整提示词
4. 在 Langfuse 中查看所有函数调用的详细追踪记录

### 验证 Docker 容器设置

确认容器配置正确：

```bash
go run cmd/ftester/main.go -flow 123 terminal -command "env | grep -i proxy" -message "检查代理设置"
```

</details>

<details>
<summary><b>在 Docker 容器中使用</b>（点击展开）</summary>

如果 PentAGI 在 Docker 中运行，可以从容器内使用 ftester：

```bash
# 在运行中的 PentAGI 容器内执行 ftester
docker exec -it pentagi /opt/pentagi/bin/ftester [arguments]

# 示例：
docker exec -it pentagi /opt/pentagi/bin/ftester -flow 123 describe
docker exec -it pentagi /opt/pentagi/bin/ftester -flow 123 terminal -command "ps aux" -message "列出进程"
```

生产部署没有本地开发环境时，这种方式很方便。

</details>

<details>
<summary><b>集成可观测性工具</b>（点击展开）</summary>

通过 ftester 发出的所有函数调用都会记录到：

1. **Langfuse**：捕获完整的 AI 智能体交互链，包括提示词、响应和函数调用
2. **OpenTelemetry**：记录指标、追踪数据和日志，用于分析系统性能
3. **终端输出**：即时反馈函数执行情况

查看详细日志：

- 在 Langfuse 界面中查看 AI 智能体追踪记录（通常位于 `http://localhost:4000`）
- 通过 Grafana 仪表板查看系统指标（通常位于 `http://localhost:3000`）
- 检查终端输出，获取即时函数结果和错误

</details>

### 命令行选项

主工具支持以下选项：

- `-env <path>`：环境文件路径（可选，默认：`.env`）
- `-provider <type>`：使用的提供商类型（默认：`custom`；可选：`openai`、`anthropic`、`ollama`、`bedrock`、`gemini`、`custom`）
- `-flow <id>`：用于测试的任务流 ID（0 表示使用模拟模式，默认：`0`）
- `-task <id>`：智能体上下文中的任务 ID（可选）
- `-subtask <id>`：智能体上下文中的子任务 ID（可选）

函数专用参数放在函数名后，并使用 `-name value` 格式传入。

### 渗透测试提示词方法论

为攻击性安全工作优化提示词时，应向智能体提供清晰的方法论，而不是简单罗列 payload（载荷）：

1. 首先明确范围、授权和成功标准
2. 先梳理应用：角色、路由、参数、上传功能、集成和信任边界
3. 系统地确定攻击面的优先级，不要一次测试所有内容
4. 先用可复现证据验证发现，再开展更深入的漏洞利用
5. 最后整理可直接用于报告的记录，写明影响、前提条件和后续步骤

PentAGI 专用提示词指南参见 [`backend/docs/prompt_engineering_pentagi.md`](backend/docs/prompt_engineering_pentagi.md)。如需从实际示例入手，可以复用并调整 [`examples/prompts/base_web_pentest.md`](examples/prompts/base_web_pentest.md)，使其符合目标应用、技术栈和评估范围。

## 构建

### 构建 Docker 镜像

Docker 构建过程会自动嵌入 Git 标签中的版本信息。请使用项目提供的脚本为构建设置正确版本：

#### Linux/macOS

```bash
# 加载版本变量
source ./scripts/version.sh

# 标准构建
docker build \
  --build-arg PACKAGE_VER=$PACKAGE_VER \
  --build-arg PACKAGE_REV=$PACKAGE_REV \
  -t pentagi:$PACKAGE_VER .

# 多平台构建
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --build-arg PACKAGE_VER=$PACKAGE_VER \
  --build-arg PACKAGE_REV=$PACKAGE_REV \
  -t pentagi:$PACKAGE_VER .

# 构建并推送
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --build-arg PACKAGE_VER=$PACKAGE_VER \
  --build-arg PACKAGE_REV=$PACKAGE_REV \
  -t myregistry/pentagi:$PACKAGE_VER \
  --push .
```

#### Windows（PowerShell）

```powershell
# 加载版本变量
. .\scripts\version.ps1

# 标准构建
docker build `
  --build-arg PACKAGE_VER=$env:PACKAGE_VER `
  --build-arg PACKAGE_REV=$env:PACKAGE_REV `
  -t pentagi:$env:PACKAGE_VER .

# 多平台构建
docker buildx build `
  --platform linux/amd64,linux/arm64 `
  --build-arg PACKAGE_VER=$env:PACKAGE_VER `
  --build-arg PACKAGE_REV=$env:PACKAGE_REV `
  -t pentagi:$env:PACKAGE_VER .
```

#### 不含版本信息的快速构建

如需构建不跟踪版本的开发镜像：

```bash
docker build -t pentagi:dev .
```

> [!NOTE]
> - 构建脚本会根据 Git 标签自动确定版本
> - 发布构建（标签所指提交）不含修订版本后缀
> - 开发构建（标签之后的提交）将提交哈希作为修订版本，例如 `1.1.0-bc6e800`
> - 如需在本地使用构建的镜像，请更新 `docker-compose.yml` 中的镜像名称，或使用 `build` 选项

## 致谢

本项目参考了以下研究与项目：
- [LLM 应用的新兴架构（Emerging Architectures for LLM Applications）](https://lilianweng.github.io/posts/2023-06-23-agent)
- [自主 LLM 智能体综述（A Survey of Autonomous LLM Agents）](https://arxiv.org/abs/2403.08299)
- Andriy Semenets 开发的 [Codel](https://github.com/semanser/codel)：为基于智能体的自动化提供了早期架构思路

## 许可证

**PentAGI** 采用 [MIT 许可证](LICENSE)。

版权所有 (c) 2025 PentAGI 开发团队

### 第三方依赖项

所有第三方依赖项均使用与 MIT 兼容的许可证。详细许可证报告参见 [licenses/](licenses/) 目录。

### VXControl 云服务

⚠️ **注意：** VXControl Cloud SDK 代码采用 MIT 许可证，但访问 **VXControl 云服务**（威胁情报、AI 支持、高级功能）需要单独的许可证密钥，并须遵守[服务条款](https://github.com/vxcontrol/cloud#license-and-terms)。

SDK 代码本身可免费使用，访问服务需要注册。

如有问题，请联系 **info@pentagi.com** 或 **info@vxcontrol.com**。

package hadoop

// namenode const
const (
	NameNodeJvmMetrics   = "Hadoop:service=NameNode,name=JvmMetrics"
	NameNodeActivity     = "Hadoop:service=NameNode,name=NameNodeActivity"
	NameNodeFSNamesystem = "Hadoop:service=NameNode,name=FSNamesystem"
	NameNodeRPC          = "Hadoop:service=NameNode,name=RpcActivityForPort"
	NameNodeStatus       = "Hadoop:service=NameNode,name=NameNodeStatus"
)

// datanode const
const (
	DataNodeJvmMetrics = "Hadoop:service=DataNode,name=JvmMetrics"
	DataNodeActivity   = "Hadoop:service=DataNode,name=DataNodeActivity"
	DataNodeRPC        = "Hadoop:service=DataNode,name=RpcActivityForPort"
)

// journalnode const
const (
	JournalNodeJvmMetrics = "Hadoop:service=JournalNode,name=JvmMetrics"
	JournalNode           = "Hadoop:service=JournalNode,name=Journal-"
	JournalNodeRPC        = "Hadoop:service=JournalNode,name=RpcActivityForPort"
)

// Metrics const
const (
	JvmMetrics = ",name=JvmMetrics"
)

// nodemanager const
const (
	NodeManagerJvmMetrics = "Hadoop:service=NodeManager,name=JvmMetrics"
	NodeManagerMetrics    = "Hadoop:service=NodeManager,name=NodeManagerMetrics"
	NodeManagerRPC        = "Hadoop:service=NodeManager,name=RpcActivityForPort"
)

// resourcemanager const
const (
	ResourceManagerJvmMetrics     = "Hadoop:service=ResourceManager,name=JvmMetrics"
	ResourceManagerQueueMetrics   = "Hadoop:service=ResourceManager,name=QueueMetrics"
	ResourceManagerClusterMetrics = "Hadoop:service=ResourceManager,name=ClusterMetrics"
	ResourceManagerRPC            = "Hadoop:service=ResourceManager,name=RpcActivityForPort"
)

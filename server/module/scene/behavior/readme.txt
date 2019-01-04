状态:
SUCCESS Status = 1
FAILURE Status = 2
RUNNING Status = 3
ERROR   Status = 4


composite
	priority //子节点只要有一个不是失败状态就返回
	memprority //只要有一个不是失败状态就返回、但是运行状态的节点还需要保留到下一个tick

	sequence     //子节点只要有一个不是成功状态就返回
	memsequence  //子节点只要有一个不是成功状态就返回  但是运行状态的节点还需要保留到下一个tick

decorator
	inverter  转换状态
	limiter   循环次数限制

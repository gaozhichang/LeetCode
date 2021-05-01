"""
【简单|难】
给定一个保存员工信息的数据结构，它包含了员工 唯一的 id ，重要度 和 直系下属的 id 。

比如，员工 1 是员工 2 的领导，员工 2 是员工 3 的领导。他们相应的重要度为 15 , 10 , 5 。那么员工 1 的数据结构是 [1, 15, [2]] ，员工 2的 数据结构是 [2, 10, [3]] ，员工 3 的数据结构是 [3, 5, []] 。注意虽然员工 3 也是员工 1 的一个下属，但是由于 并不是直系 下属，因此没有体现在员工 1 的数据结构中。

现在输入一个公司的所有员工信息，以及单个员工 id ，返回这个员工和他所有下属的重要度之和。

75%，30%
解题：
方法一：深度优先搜索
深度优先搜索的做法非常直观。根据给定的员工编号找到员工，从该员工开始遍历，对于每个员工，将其重要性加到总和中，然后对该员工的每个直系下属继续遍历，直到所有下属遍历完毕，此时的总和即为给定的员工及其所有下属的重要性之和。

实现方面，由于给定的是员工编号，且每个员工的编号都不相同，因此可以使用哈希表存储每个员工编号和对应的员工，即可通过员工编号得到对应的员工。


卡点：
1. 先把员工编程k:v的字典，再进行父级查找，不会有先出现儿子，再出现爸爸的问题。
2. 这里直接对字典（从目标id出发），进行dfs深度遍历。巧妙。我一直用原数组

"""

"""
# Definition for Employee.
class Employee(object):
    def __init__(self, id, importance, subordinates):
    	#################
        :type id: int
        :type importance: int
        :type subordinates: List[int]
        #################
        self.id = id
        self.importance = importance
        self.subordinates = subordinates
"""

class Solution(object):
    def getImportance(self, employees, id):
        """
        :type employees: List[Employee]
        :type id: int
        :rtype: int
        """
        eid_obj = {}
        for i in employees:
            eid_obj[i.id] = i
        
        return self.dfs(eid_obj,id)

    def dfs(self,eid_obj,id):
        total = 0
        if eid_obj.has_key(id):
            total = eid_obj[id].importance
            for i in eid_obj[id].subordinates:
                total = self.dfs(eid_obj,i)+total

        return total

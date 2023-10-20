
class Node:

    def __init__(self, data=None , next_node=None ):
        self.data  = data
        self.next_node = next_node


class LinkedList:

    def __init__(self):
        
        self.head = None
        self.last_node = None

    
    def to_array(self):
        arr = []

        if self.head is None:
            return arr
        

        node = self.head

        while node:
            arr.append(node.data)
            node = node.next_node

        return arr


    def print_linked_list(self):
        
        linked_list_string = ""
        node = self.head

        while node:
            linked_list_string += f"{node.data} ->"
            node = node.next_node

        linked_list_string += "None"

        print(linked_list_string)



    def insert_at_head(self,data):
    
        if self.head is None:               
            self.head = Node(data , None)
            self.last_node = self.head                       # optimize for insert_at_tail() as well as that can be then done without the while loop !
            # return 
        
        new_node = Node(data , self.head)
        # new_node.next_node = self.head
        self.head = new_node


    def insert_from_tail(self,data):

        if self.head is None:
            self.insert_at_head(data)
            return 

        new_node = Node(data , None)
        node = self.head

        if self.last_node is None:                                    #commented out as if we user insert_at_beginning() , we keep tract of the last node as well , so we dont really need this if block of code !

            while node.next_node:
                node = node.next_node

            node.next_node = Node(data , None)
            self.last_node = node.next_node

        else:                                                       # with this 0(n) tail insert will only happen for the first time , from then onwards it will be o(1) as we know the position of the last node!
        
            self.last_node.next_node = Node(data , None)
            self.last_node = self.last_node.next_node
  
    
    def insert_between(self, data , value):

        node = self.head
        new_node = Node(data , None)

        if self.head is None:
            self.insert_at_head(data)
            return 
        
        while node:

            if node.data == value:
                new_node.next_node = node.next_node
                node.next_node = new_node
                return

            node = node.next_node

        return "Value does not exist in LinkedList"


    def clear(self):
        self.head = None
        self.last_node = None


    def delete_from_head(self):
        
        if self.head is None:
            return "No items in LinkedList"
        
        self.head = self.head.next_node


    def delete_from_tail(self):
        
        node = self.head

        if self.head is None:
            return "Empty linkedList"

        if node.next_node is None:
            self.delete_from_head()

        while node.next_node.next_node:
            node = node.next_node
        
        node.next_node = None


    def delete_by_value(self , value):

        if self.head == None:
            return "Empty LinkedList"

        if value == self.head.data:
            self.delete_from_head()
            return 

        node = self.head

        while node.next_node:


            if node.next_node.data == value:
                node.next_node = node.next_node.next_node
                return
            
            node = node.next_node


    def search(self , value):

        node = self.head
        position = 0

        while node:
            if node.data == value:
                return f"Found the node in location {position} !"
            node = node.next_node
            position += 1

        return "not found"

    def search_by_index(self , index):                      # use __getitem__() to use array [] syntax for finding the index values 
        
        node = self.head
        position = 0

        while node:

            if index == position:
                return node.data
            
            node = node.next_node
            position += 1

        return "IndexError"
    

    def replace_max_num_with_given_value(self , value):


        node = self.head
        max_node = node

        while node:
            if node.data > max_node.data:
                max_node = node
            node = node.next_node

        max_node.data = value

        

        


ll = LinkedList()

n4 = Node(4 , None)
n3 = Node(3 , n4)
n2 = Node(190 , n3)
n1 = Node(1 , n2)

ll.head = n1

# ll.delete_by_value(1)
# ll.delete_by_value(2)
# ll.delete_from_tail()
# ll.delete_by_value(3)
# ll.print_linked_list()


ll.print_linked_list()

class Game:
    def __init__(self, name, year, sales):
        self.name = name
        self.year = year
        self.sales = sales

    def getName(self):
        return self.name

    def getYear(self):
        return self.year

    def getSales(self):
        return self.sales

def compareItems(itemA, itemB, reverse, flag):
    if flag == 'name' and reverse:
        return itemA.getName() > itemB.getName()
    elif flag == 'year' and reverse:
        return itemA.getYear() > itemB.getYear()
    elif flag == 'sales' and reverse:
        return itemA.getSales() > itemB.getSales()
    elif flag == 'name':
        return itemA.getName() < itemB.getName()
    elif flag == 'year':
        return itemA.getYear() < itemB.getYear()
    elif flag == 'sales':
        return itemA.getSales() < itemB.getSales()
    else:
        print('Command not found')
        return False

def mergeSort(arr, reverse, flag):
    if len(arr) > 1:
        mid = len(arr)//2
        left = arr[:mid]
        right = arr[mid:]
        mergeSort(left,reverse, flag)
        mergeSort(right,reverse, flag)
        i,j,k = (0, 0, 0)

        while i < len(left) and j < len(right):
            if compareItems(left[i], right[j], reverse, flag):
                arr[k] = left[i]
                i += 1
            else:
                arr[k] = right[j]
                j += 1

            k += 1

        while i < len(left):
            arr[k] = left[i]
            i += 1
            k += 1
 
        while j < len(right):
            arr[k] = right[j]
            j += 1
            k += 1

alist = [ Game('LOL', 2013, 0), Game('PES', 2021, 10000), Game('AGARIO', 2010, 1) ]
mergeSort(alist,False,'year')
for i in alist:
    print(i.getName())

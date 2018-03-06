# Rain Trap

# Given an array consisting of n non-negative integers representing elevation, compute how much water it is able to trap after raining.

#  5  +
#     |
#  4  +
#     |
#  3  +                           +---+
#     |                           |   |
#  2  +           +---+           +-------+   +---+
#     |           |   | x   x   x |   |   | x |   |
#  1  +   +---+   +-------+   +-----------------------+
#     |   |   | x |   |   | x |   |   |   |   |   |   |
#     +-----------------------------------------------+
#       0   1   0   2   1   0   1   3   2   1   2   1 

# The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rainwater (labeled with an x) are being trapped.



def calcWater(elevations):
	eStack = []
	waterUnits = 0
	last = None
	maxElevation = max(elevations)
	total = 0

	for j in xrange(maxElevation):
		j += 1
		last = -1
		for index, elev in enumerate(elevations):
			if last > -1 and elev >= j:
				total += index - last - 1
			if elev >= j:
				last = index
	return total
	


def robberHelper(values, skippedFirst):
	#print(values)
	#print(skippedFirst)
	if len(values) == 0 or (len(values) == 1 and not skippedFirst):
		return 0
	if len(values) == 1:
		return values[0]
	return max(values[0] + robberHelper(values[2:], skippedFirst), robberHelper(values[1:], skippedFirst) )


def robber(values):
	if len(values) == 0:
		return 0
	#print(values)
	return max(values[0]+robberHelper(values[2:], False), robberHelper(values[1:], True), robberHelper(values[2:], True))

# 
# [1,2,3,4,5,5]
# 






x = [5,3,0,3]
y = [4,0,0,4]
z = [0,1,0]
print(robber(x))
#print(calcWater(z))



	
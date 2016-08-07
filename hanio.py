#this is my python hanio
def hanio(a,b,c,n):
	if n==1:
		print a+"=>"+c
		return 
	hanio(a,c,b,n-1)
	print a+"=>"+c
	hanio(b,a,c,n-1)
hanio("a","b","c",3)
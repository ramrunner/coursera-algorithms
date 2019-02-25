all:
	javac -cp ../../algs4.jar:. WordNet.java
	javac -cp ../../algs4.jar:. SAP.java
	javac -cp ../../algs4.jar:. Outcast.java

test-wordnet:
	java -cp ../../algs4.jar:. WordNet ../synsets.txt ../hypernyms.txt

test-sap:
	java -cp ../../algs4.jar:. SAP ../digraph1.txt

test-outcast:
	java -cp ../../algs4.jar:. Outcast ../synsets.txt ../hypernyms.txt ../outcast5.txt ../outcast8.txt ../outcast11.txt

clean:
	rm *.class

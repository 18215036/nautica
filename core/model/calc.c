//Dive Table Calculator (based on US DIVING MANUAL REV.7
//created by: Ahmad Fawwaz Zuhdi

//NO DECO DIVE HANDLING
//Add another dive (y/n)?Add another dive (y/n)? //problem?
//Adding save option later

#include <stdio.h>
#include <math.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>

typedef struct{
	int D;	//depth
	int R; 	//residual
	int A;	//actual
	int T;	//total
	int S;	//surface
	char P;	//previous group
	char C;	//current group
} dive;

typedef struct{
	int A;
	int B;
	int C;
	int D;
	int E;
	int F;
	int G;
	int H;
	int I;
	int J;
	int K;
	int L;
	int M;
	int N;
	int O;
	int Z;
} DATA2;		//surface interval data

typedef struct{
	DATA2 d2;
	int DE;		//depth
} DATA3;		//residual nitrogen data

typedef struct{
	DATA3 d3;
	int NSL;	//no-stop limit
} DATA1;		//no-deco data

DATA1 t1[24];
DATA2 t2[16];
DATA3 t3[24];

//dataOne Initializer
void dataOne(){
	FILE *fp = fopen("D1", "r");
	
	if(fp != NULL){
		char line[68];
		char *token;
		const char *s = " ,";		//delimiter
		int i=0;
		
		while(fgets(line, sizeof line, fp) != NULL){
			token = strtok(line, s);
			t1[i].d3.DE = atoi(token);
			token = strtok(NULL,s);
			t1[i].NSL = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.A = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.B = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.C = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.D = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.E = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.F = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.G = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.H = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.I = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.J = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.K = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.L = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.M = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.N = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.O = atoi(token);
			token = strtok(NULL,s);
			t1[i].d3.d2.Z = atoi(token);
			i++;
		}
		fclose(fp);
	} else {
		perror("D1");
	}
}

//dataTwo Initializer
void dataTwo(){
	FILE *fp = fopen("D2", "r");
	
	if(fp != NULL){
		char line[64];
		char *token;
		const char *s = " ,";		//delimiter
		int i=0;
		
		while(fgets(line, sizeof line, fp) != NULL){
			token = strtok(line, s);
			t2[i].A = atoi(token);
			token = strtok(NULL,s);
			t2[i].B = atoi(token);
			token = strtok(NULL,s);
			t2[i].C = atoi(token);
			token = strtok(NULL,s);
			t2[i].D = atoi(token);
			token = strtok(NULL,s);
			t2[i].E = atoi(token);
			token = strtok(NULL,s);
			t2[i].F = atoi(token);
			token = strtok(NULL,s);
			t2[i].G = atoi(token);
			token = strtok(NULL,s);
			t2[i].H = atoi(token);
			token = strtok(NULL,s);
			t2[i].I = atoi(token);
			token = strtok(NULL,s);
			t2[i].J = atoi(token);
			token = strtok(NULL,s);
			t2[i].K = atoi(token);
			token = strtok(NULL,s);
			t2[i].L = atoi(token);
			token = strtok(NULL,s);
			t2[i].M = atoi(token);
			token = strtok(NULL,s);
			t2[i].N = atoi(token);
			token = strtok(NULL,s);
			t2[i].O = atoi(token);
			token = strtok(NULL,s);
			t2[i].Z = atoi(token);
			i++;
		}
		fclose(fp);
	} else {
		perror("D2");
	}
}
	
//dataThree Initializer
void dataThree(){
	FILE *fp = fopen("D3", "r");
	
	if(fp != NULL){
		char line[61];
		char *token;
		const char *s = " ,";		//delimiter
		int i=0;
		
		while(fgets(line, sizeof line, fp) != NULL){
			token = strtok(line, s);
			t3[i].DE = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.Z = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.O = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.N = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.M = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.L = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.K = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.J = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.I = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.H = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.G = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.F = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.E = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.D = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.C = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.B = atoi(token);
			token = strtok(NULL,s);
			t3[i].d2.A = atoi(token);
			i++;
		}
		fclose(fp);
	} else {
		perror("D3");
	}
}

//program Intiliazer
void start(){
	dataOne();
	dataTwo();
	dataThree();
}

//convert meter to feet
int mToF(int m){
	return ceil(3.28084*m);
}

//convert group number to group letter
char nToL(int n){
	switch (n)
	{
		case 1:return 'A';
		case 2:return 'B';
		case 3:return 'C';
		case 4:return 'D';
		case 5:return 'E';
		case 6:return 'F';
		case 7:return 'G';
		case 8:return 'H';
		case 9:return 'I';
		case 10:return 'J';
		case 11:return 'K';
		case 12:return 'L';
		case 13:return 'M';
		case 14:return 'N';
		case 15:return 'O';
		case 16:return 'Z';
		default:return '0';	
	}
}

//TABLE 1 FUNCTION //-6 = deco dive
char f1(int depth, int total){
	int i = 0;
	int ft = mToF(depth);
	
	while (ft > t1[i].d3.DE)
	{
		i++;
	}
	
	return (total > t1[i].NSL) ? -6 :
			((total <= t1[i].d3.d2.A) ? 'A' :
			((total <= t1[i].d3.d2.B) ? 'B' :
			((total <= t1[i].d3.d2.C) ? 'C' :
			((total <= t1[i].d3.d2.D) ? 'D' :
			((total <= t1[i].d3.d2.E) ? 'E' :
			((total <= t1[i].d3.d2.F) ? 'F' :
			((total <= t1[i].d3.d2.G) ? 'G' :
			((total <= t1[i].d3.d2.H) ? 'H' :
			((total <= t1[i].d3.d2.I) ? 'I' :
			((total <= t1[i].d3.d2.J) ? 'J' :
			((total <= t1[i].d3.d2.K) ? 'K' :
			((total <= t1[i].d3.d2.L) ? 'L' :
			((total <= t1[i].d3.d2.M) ? 'M' :
			((total <= t1[i].d3.d2.N) ? 'N' :
			((total <= t1[i].d3.d2.O) ?	'O' :
										'Z')))))))))))))));
}

//TABLE 2 FUNCTION //-4 = not a repetitive dive //-5 = SI<10
char f2(char pGroup, int S){
	int i = 0;
	
	while (nToL(i+1) != pGroup)
	{
		i++;
	}
	
	return (S <= t2[i].A) ?
			((S <= t2[i].B) ?
			((S <= t2[i].C) ?
			((S <= t2[i].D) ?
			((S <= t2[i].E) ?
			((S <= t2[i].F) ?
			((S <= t2[i].G) ?
			((S <= t2[i].H) ?
			((S <= t2[i].I) ?
			((S <= t2[i].J) ?
			((S <= t2[i].K) ?
			((S <= t2[i].L) ? 
			((S <= t2[i].M) ?
			((S <= t2[i].N) ?
			((S <= t2[i].O) ?
			((S <= t2[i].Z) ?
			((S < 10) ? -5 : 'Z') : 'O') : 'N') : 'M') : 'L') : 'K') : 'J') : 'I') : 'H') : 'G') : 'F') : 'E') : 'D') : 'C') : 'B') : 'A') : -4;
}

//TABLE 3 FUNCTION //-2? //-3?
int f3(char pGroup, int nDepth){
	int i = 0;
	int ft = mToF(nDepth);
	
	while (ft > t3[i].DE)
	{
		i++;
	}
	
	switch (pGroup)
	{
		case 'A': return t3[i].d2.A;
		case 'B': return t3[i].d2.B;
		case 'C': return t3[i].d2.C;
		case 'D': return t3[i].d2.D;
		case 'E': return t3[i].d2.E;
		case 'F': return t3[i].d2.F;
		case 'G': return t3[i].d2.G;
		case 'H': return t3[i].d2.H;
		case 'I': return t3[i].d2.I;
		case 'J': return t3[i].d2.J;
		case 'K': return t3[i].d2.K;
		case 'L': return t3[i].d2.L;
		case 'M': return t3[i].d2.M;
		case 'N': return t3[i].d2.N;
		case 'O': return t3[i].d2.O;
		case 'Z': return t3[i].d2.Z;
		default: return 0;
	}
}

dive d[5];

//print Dives
void printDives(int nDives){
	int i;
	
	for (i = 0; i < nDives; i++)
	{
		printf("DIVE %d\n",i+1);
		printf("PPG: %c\n",d[i].P);
		printf("Depth: %d\n",d[i].D);
		printf("RNT: %d\n",d[i].R);
		printf("ABT: %d\n",d[i].A);
		printf("TBT: %d\n",d[i].T);
		printf("CPG: %c\n",d[i].C);
		printf("SI: %d\n",d[i].S);
		printf("\n");
	}
}

int main(int argc, char **argv)
{
	int i = 0;
	int depth;
	int actual;
	int surface;
	bool firstDive = true;
	bool decoDive = false;
	char c = 'y';
	
	start();
	while(!decoDive && i<5 && c=='y')
	{
		c = 0;
		printf("DIVE %d\n",i+1);
		printf("Depth: ");scanf("%d",&depth);
		d[i].D = depth;
		printf("Actual: ");scanf("%d",&actual);
		d[i].A = actual;
		printf("Surface: ");scanf("%d",&surface);
		d[i].S = surface;
		
		if (firstDive)
		{
			d[i].P = 0;
			d[i].R = 0;
			firstDive = !firstDive;
		}
		else
		{
			d[i].P = f2(d[i-1].C,d[i-1].S);
			if (d[i].P == -4)
			{
				d[i].P = 0;
				d[i].R = 0;
			}
			else
			{
				if (d[i].P == -5)
				{
					d[i].D = (d[i].D > d[i-1].D) ? d[i].D : d[i-1].D;
					d[i].A += d[i-1].A;
					d[i].S += d[i-1].S;	//added Surface Interval?
					d[i].P = d[i-1].P;
				}
				d[i].R = f3(d[i].P,d[i].D);
			}
		}
		
		if (d[i].R == -2)
		{
			d[i].T = d[i].A;
			d[i].C = d[i].P;
		}
		else if (d[i].R == -3)
		{
			d[i].T = -6; //deco
			d[i].C = -6;
		}
		else
		{
			d[i].T = d[i].R + d[i].A;
			d[i].C = f1(d[i].D,d[i].T);
		}
		printDives(i+1);
		
		decoDive = (d[i].C == -6) ? true : false;
		while (!(c=='y' || c=='n'))
		{
			printf("Add another dive (y/n)?");
			scanf("%c",&c);
		}
		i++;
	}
	
	(decoDive) ? printf("DIVE %d is a DECO DIVE\n",i) : printf("\n");
	printDives(i);
	
	return 0;
}

/* -1 : Unlimited No-Stop Limit */
/* -2 : Residual Nitrogen Time cannot be determined using this table (see paragraph 9-9.1 subparagraph 8 for instructions).
 * At depths of 10, 15, and 20 fsw, some of the higher repetitive groups do not
have a defined residual nitrogen time. These groups are marked with a double
asterisk in the lower half of Table 9-8. The RNT is undefined because the tissue
nitrogen loading associated with those repetitive groups is higher than the
nitrogen loading that could be achieved even if the diver were to remain at
those depths for an infinite period of time. A diver entering the dive in one of
those higher groups marked by a double asterisk can still perform a repetitive
dive at 10, 15 or 20 fsw because the no-decompression time at those depths
is unlimited. An RNT time is not required to make the dive. If a subsequent
repetitive dive to a deeper depth is planned, however, the diver will need a
repetitive group at the end of the shallow dive in order to continue using the
RNT table. If a double asterisk is encountered in Table 9-8, assume that the
repetitive group remains unchanged during the course of the dive at 10, 15, or
20 fsw.*/
/* -3 : Read vertically downward to the 30 fsw repetitive dive depth. Use the corresponding residual nitrogen times to compute the
equivalent single dive time. Decompress using the 30 fsw air decompression table. */

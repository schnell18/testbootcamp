all:    quiz

quiz: quiz.tex
	xelatex -shell-escape -synctex=1 quiz
	xelatex -shell-escape -synctex=1 quiz

answer: quiz.tex
	cat quiz.tex | sed -e '/^\\documentclass/ s/\[\(.*\)\]/[\1,answers]/' | xelatex -shell-escape -synctex=1
	cat quiz.tex | sed -e '/^\\documentclass/ s/\[\(.*\)\]/[\1,answers]/' | xelatex -shell-escape -synctex=1

clean:
	rm -fr _minted* *.aux *.log *.nav *.out *.pdf *.snm *.toc *.vrb *.pyg

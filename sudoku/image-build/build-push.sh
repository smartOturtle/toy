CGO_ENABLED=0 go build .
docker build -t dockeridfor823/sudoku:latest .
# docker run  -p 1234:1234 dockeridfor823/sudoku:latest
docker push dockeridfor823/sudoku:latest
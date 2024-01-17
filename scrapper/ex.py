import csv

# CSV 파일 경로
file_path = 'all-weeks-countries.tsv'
write = "countries.txt"
# 첫 번째 열의 데이터를 저장할 집합 (Set)
first_column_values = set()

# CSV 파일 열기 및 읽기
with open(file_path, mode='r', encoding='utf-8') as file:
    reader = csv.reader(file,delimiter="\t")
    for row in reader:
        if row:  # 행이 비어있지 않은 경우에만 처리
            first_column_values.add(row[0])  # 첫 번째 열의 값 추가

# 결과 출력
#first_column_values.remove("country_name")
print(first_column_values)

with open(write, mode = "w", encoding="utf-8") as file:
    for v in first_column_values:
        v = v.lower()
        v = v.replace(' ','-')
        if v == "country_name":
            continue
        file.write(v+"\n")

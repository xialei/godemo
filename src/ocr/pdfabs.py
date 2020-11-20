from pdf_layout_scanner import layout_scanner
toc = layout_scanner.get_toc('/Users/roger/workspace/go/src/files/yxkj.pdf')
print(len(toc))

print(toc)


pages = layout_scanner.get_pages('/Users/roger/workspace/go/src/files/yxkj.pdf')
print(pages[0])
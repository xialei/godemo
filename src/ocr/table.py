from PIL import Image, ImageOps
import cv2
import numpy as np

def get_points(img_transverse, img_vertical):
    """
    获取横纵线的交点
    :param img_transverse:
    :param img_vertical:
    :return:
    """
    img = cv2.bitwise_and(img_transverse, img_vertical)
    return img

def dilate_img(img, kernal_args:tuple, iterations:int):
    """
    dilate image
    @param kernel_args 卷积核参数（2，2）
    @param interations dilate的迭代次数
    """
    kernel = np.ones(kernal_args, np.uint8)
    return cv2.dilate(img, kernel,iterations=iterations)

def erode_img(img,kernel_args=(2,2),iterations=1, d='h'):
    """
    对图像进行腐蚀
    @param kernel_args 卷积核参数（2，2）
    @param interations erode的迭代次数
    """
    kernel = np.ones(kernel_args, np.uint8)
    return cv2.erode(img, kernel,iterations=iterations)

def bin_img(img:'numpy.ndarray'):
    """
    对图像进行二值化处理
    :param img: 传入的图像对象（numpy.ndarray类型）
    :return: 二值化后的图像
    """
    ret,binImage=cv2.threshold(img,220,255,cv2.THRESH_BINARY_INV)
    cv2.imwrite("/Users/roger/workspace/output/test_bin2.png", binImage)
    return binImage

def gray_img(img:'numpy.ndarray'):
    """
    对读取的图像进行灰度化处理
    :param img: 通过cv2.imread(imgPath)读取的图像数组对象
    :return: 灰度化的图像
    """
    grayImage=cv2.cvtColor(img,cv2.COLOR_BGR2GRAY)
    return grayImage

def split_rec(arr):
    """
    切分单元格
    :param arr:
    :return:
    """
    # 数组进行排序
    arr.sort(key=lambda x: x[0],reverse=True)
    # 数组反转
    arr.reverse()
    for i in range(len(arr) - 1):
        # print("==========", i)
        if arr[i+1][0] == arr[i][0]:
            arr[i+1][3] = arr[i][1]
            arr[i+1][2] = arr[i][2]
        if arr[i+1][0] > arr[i][0]:
            arr[i+1][2] = arr[i][0]
        # print(arr[i])

    return arr

def get_rec(img):
    """
    获取单元格
    :param img:
    :return:
    """
    # 检测轮廓，矩形4个顶点
    contours, hierarchy = cv2.findContours(img, cv2.RETR_CCOMP, cv2.CHAIN_APPROX_SIMPLE)
    contours_poly = [0] * len(contours)
    boundRect = [0] * len(contours)
    rois = []
    print(len(contours))
    # for i in range(len(contours) - 1):
    for i in range(len(contours)):
        cnt = contours[i]
        # 把一个连续光滑曲线折线化,对图像轮廓点进行多边形拟合
        contours_poly[i] = cv2.approxPolyDP(cnt, 1, True)
        # 计算轮廓的垂直边界最小矩形
        boundRect[i] = cv2.boundingRect(contours_poly[i])
        rois.append(np.array(boundRect[i]))
        # img = cv2.rectangle(img_bak, (boundRect[i][0], boundRect[i][1]), (boundRect[i][2], boundRect[i][3]),
        #                     (255, 255, 255), 1, 8, 0)
    rois = split_rec(rois)
    return rois

if __name__ == "__main__":
    image  = "/Users/roger/workspace/test.png"
    img_bak = cv2.imread(image)
    img = gray_img(img_bak)
    img = bin_img(img)
    
    # 纵向腐蚀(变瘦去噪声)获取横向线条 ====
    img_transverse = erode_img(img, (1,2), 30)
    cv2.imwrite("/Users/roger/workspace/output/test_ero_h.png", img_transverse)
    # 横向腐蚀获取纵向线条 ||||
    img_vertical = erode_img(img, (2,1), 30)
    cv2.imwrite("/Users/roger/workspace/output/test_ero_v.png", img_vertical)
    
    # 膨胀处理对线条进行加粗
    img_transverse = dilate_img(img_transverse,(2,2),1)
    cv2.imwrite("/Users/roger/workspace/output/test_dilate_h.png", img_transverse)
    img_vertical = dilate_img(img_vertical,(2,2),1)
    cv2.imwrite("/Users/roger/workspace/output/test_dilate_v.png", img_vertical)
    
    # 获取交点 :::::
    img = get_points(img_transverse,img_vertical)
    cv2.imwrite("/Users/roger/workspace/output/test_point.png", img)

    rois = get_rec(img)
    for i, r in enumerate(rois):
        # cv2.imshow("src" + str(i), img_bak[r[3]:r[1], r[2]:r[0]])
        cv2.imwrite("/Users/roger/workspace/output/cell/cell_"+str(i)+".png", img_bak[r[3]:r[1], r[2]:r[0]])
    # cv2.waitKey(0)

    # cv2.destroyAllWindows()
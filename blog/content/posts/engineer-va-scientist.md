---
title: "Engineer và scientist"
date: 2018-04-29T17:48:21+07:00
draft: false
description: "Nhìn thấy lớp trẻ bây giờ thật có ý chí tiến thủ, chả bù cho một thanh niên đi làm vài năm cũng ở tầm gọi là lão làng rồi nhưng vẫn chỉ là thằng engineer quèn mình cảm thấy chạnh lòng thực sự. Mình hỏi thêm là lý do tại sao em lại muốn làm scientist, ngoài nó ra thì vẫn có data engineer mà? Bạn trả lời bạn ấy muốn làm data scientist vì... nghề đó đang hot và nghe cool hơn dăm ba mấy thằng engineer nhiều lắm."
tags: [engineer, random, scientist]
---

“I am not a visionary. I'm an engineer. I'm happy with the people who are wandering around looking at the stars but I am looking at the ground and I want to fix the pothole before I fall in.”
― Linus Torvalds

#### Chuyện ngày hôm ấy
Đợt vừa rồi lúc còn làm việc ở công ty C nọ, vô tình mình được một dịp phỏng vấn một bạn fresher apply vào một vị trí của team data. Nói thật là mình chẳng có một tí kinh nghiệm nào về làm data cả nên được rủ thì mình đi theo... cho vui vậy thôi.

Sau màn chào hỏi và giới thiệu bản thân và những gì bạn ứng viên học tập được ở trong trường cũng như về project đề tài tốt nghiệp của bạn ấy. Bạn làm data quay sang hỏi tại sao chọn thuật toán xử lý đó rồi thì train data model thế nào và những thứ cao siêu tầm vũ trụ, thành thật mà nói thì mình  nghe... không hiểu.

Sau đó đến lượt mình thì mình có quay sang hỏi bạn ấy một vài vấn đề xã hội. Đầu tiên mình hỏi một câu có lẽ khá quen thuộc đó là tại sao bạn lại chọn công ty mình mà không chọn công ty khác? Câu trả lời đại loại là bạn muốn có một môi trường để tiếp cận nhiều hơn với big data và muốn sau vài năm thì lên làm data scientist.

Nhìn thấy lớp trẻ bây giờ thật có ý chí tiến thủ, chả bù cho một thanh niên đi làm vài năm cũng ở tầm gọi là lão làng rồi nhưng vẫn chỉ là thằng engineer quèn mình cảm thấy chạnh lòng thực sự. Mình hỏi thêm là lý do tại sao em lại muốn làm scientist, ngoài nó ra thì vẫn có data engineer mà? Bạn trả lời bạn ấy muốn làm data scientist vì... nghề đó đang hot và nghe cool hơn dăm ba mấy thằng engineer nhiều lắm.

Nghe đến đây thì mình cũng cạn lời, chọn một nghề chỉ bởi vì nó hot và nghe nó cool hơn, nhưng thôi mỗi người có một lý tưởng riêng, không thể lấy lý tưởng của một thằng engineer với ước mơ về quê trồng rau của mình để áp đặt cho lớp trẻ bây giờ được. Thôi thì làm một bài test về khả năng code với tư duy của bạn vậy.

#### Đề bài

Cho 2 số integer x, y, viết hàm tính kết quả phép nhân của hai số mà không dùng toán tử nhân(*). Bạn có thể dùng bất cứ ngôn ngữ nào mà bạn cảm thấy thuận tiện nhất.

#### Viết chạy được cái đã

Sau một bạn ngồi trầm ngâm suy nghĩ về nước Mỹ cũng như nghề data scientist của bạn, và với tốc độ download rùa bò 5KB/s của nhà cung cấp dịch vụ mạng Cá Mập Trắng làm tôi không thể xem được youtube nên tôi đành phải hint cho nam thanh niêm một xíu.

Hãy nghĩ đơn giản một xíu, ví dụ như một phép nhân 5 * 3 thì em có thể viết thành: 5 + 5 + 5 mà. Sau đó bạn ấy dần hiểu ra và viết cho tôi một chương trình đại loại thế này (tôi sẽ implements lại bằng code C)

```
int multi(int x, int y) {
    int sum = 0;
    for (int i=0; i<y; i++) {
        sum += x;
    }
    return sum;
}
```

#### Hỏi xoáy lần một

Hỏi: Chương trình của em thì chạy được đấy, nhưng bây giờ anh gọi hàm của em như vầy `multi(1, 1000);` em nghĩ nó mất bao lâu để trả về kết quả?

Đáp: (suy nghĩ một lúc), em sẽ kiểm x, y xem cái nào nhỏ hơn, swap giá trị rồi mới thực hiện for loop theo giá trị bé.

Hỏi: Ý tưởng hay đấy, nhưng nếu anh gọi `multi(1000, 1000);` thì sao?

Đáp: ... (lại trầm ngâm suy nghĩ)

Hỏi: Anh sẽ cho em một ý tưởng, ví dụ `1000 * 1000` nó sẽ bằng `2000 * 500` và bằng với `4000 * 250`... Và anh hint thêm một tí, tương tự `1000 * 1001` nó sẽ bằng với `2000 * 500 + 1000`

Đáp: ...(trầm ngâm về nước Mỹ tiếp)

Và đây thật sự là kết quả tôi mong muốn

```
int multi(int x, int y) {
    return y? (y > 1? multi(x << 1, y >> 1) + multi(x, y & 1): x): 0;
}
```

Phiên bản không hack não (các phần tiếp theo sẽ dùng phiên bản hack não cho... gọn):

```
int multi(int x, int y) {
    if (y == 0) {
        return 0;
    }
    if (y == 1) {
        return x;
    }
    return multi(x << 1, y >> 1) + multi(x, y & 1);
}
```

#### Hỏi xoáy lần hai

Hỏi: Chương trình như vậy là chạy được rồi đó, nhưng vẫn không pass hết mọi trường hợp, giả sử bây giờ anh gọi hàm `multi(3, -5)`, kết quả mong muốn của anh là -15 cơ, nhưng trường hợp này thì nó trả về sai mất rồi.

Đáp: ...

Hỏi: Nếu anh lấy trị tuyệt đối thì sao nhỉ?

Đáp: Lúc đó em sẽ xét dấu kết quả. Chương trình trở thành

```
int multi(int x, int y) {
    
    int sign = 0;

    if (x > 0) {
        if (y > 0) {
            sign = 0;
        } else {
            sign = 1;
        }
    } else {
        if (y < 0) {
            sign = 0;
        } else {
            sign = 1;
        }
    }

    x = abs(x);
    y = abs(y);

    int result = y? (y > 1? multi(x << 1, y >> 1) + multi(x, y & 1): x): 0;
    if (sign) {
        return -result;
    }
    return result;
}
```

#### Hỏi xoáy lần ba

Hỏi: Ý tưởng xét dấu thì đúng rồi đó, nhưng mà với cách làm này em gọi hàm multi khá nhiều lần, sẽ bị overhead chuyện đi xét dấu, đó là chưa kể dùng if else khá nhiều thế này code nhìn vào sẽ khá dơ. Em có ý tưởng gì không?

Đáp: ...(trầm ngâm)

Hỏi: Theo em thì số âm được biểu diễn như thế nào trong bộ nhớ máy tính, cụ thể là biểu diễn ở dạng bits.

Đáp: ...

Hỏi: Thật ra có rất nhiều cách biểu diễn số nguyên âm, nhưng tất cả những cách mà anh biết chung quy lại đều có một điểm chung là bit đầu tiên sẽ là sign bit, chi tiết em có thể về tìm tài liệu đọc thêm. Để xét dấu trong trường hợp của mình thì có một cách đơn giản là xor 2 số x, y rồi kiểm tra sign bit thôi.

Kết quả mà tôi mong muốn nó như thế này:

```
#include <stdlib.h>

int __multi(int x, int y) {
    return y? (y > 1? __multi(x << 1, y >> 1) + __multi(x, y & 1): x): 0;
}

int multi(int x, int y) {
    if ((x ^ y) >> 31) {
        return - __multi(abs(x), abs(y));
    }
    return __multi(abs(x), abs(y));
}
```

À mà, về cơ bản tôi chỉ là chàng trai(engineer) tử tế.

#### Bài viết này có giúp tôi hết thất nghiệp hay không?

Thành thực mà nói thì câu này là tôi mượn của anh blog [Quần Cam](https://quan-cam.com), tất nhiên là có trả tiền bản quyền đàng hoàng bằng gói trà sao suốt 35k/lạng. Trong những chuỗi ngày thất nghiệp tụt quần ngẫm nghĩ về tương lai vô định của một thằng engineer ấp ủ ước mơ được về quê trồng rau như tôi mới rảnh rỗi sinh nông nổi mà ngồi viết blog chia sẻ vài cảm nghĩ linh tinh.

Và thành thực mà nói lần nữa thì làm nghề gì không quan trọng, Chẳng phải khi bạn còn học lớp 1 cô giáo của bạn đã từng dạy là nghề nghiệp nào cũng cao quý như nhau cả, bạn là anh quét rác, Worker hay Engineer cũng vậy mà Scentist cũng như nhau, với bất cứ nghề nào nếu như bạn hiểu rõ thực sự cái bạn làm là gì thì đều có nét đẹp của riêng nó cả.

Thân ái!

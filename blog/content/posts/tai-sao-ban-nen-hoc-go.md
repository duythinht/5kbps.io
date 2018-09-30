---
title: "Tại sao bạn nên học code Go?"
date: 2018-05-14T09:44:06+07:00
draft: false 
description: "Bắt đầu từ những vấn đề thực tế, Go được tạo ra để giải quyết những bài toán đến từ real world, tập trung nhiều vào software engineering hơn là research programming language design, nhiều người sẽ cảm thấy vui vẻ để sử dụng nó như một công cụ mạnh mẽ, nhưng số khác thì cho rằng nó là một boring language, không sáng tạo, cú pháp dài dòng, xấu xí."
images: ["https://cdn0.tnwcdn.com/wp-content/blogs.dir/1/files/2018/07/go.png"]
tags: [engineer, go]
---

 Tôi có thằng bạn, thời đi học lúc tôi học Java thì nó học C#, sau ra đường kiếm sống tôi xài Python thì nó nhập hội Ruby. Gần đây thấy bảo nó lại âm thầm dụ dỗ người ta về "Hội thánh đức chúa [Elixir](https://quan-cam.com/posts/elixir-erlang-actors-model-va-concurrency)". Mà thực ra tôi nói vòng vo vậy thôi chứ mục đích của tôi khi viết bài này là **dụ mọi người học Go**.

### Go là cái gì?

> Ngôn ngữ lập trình hiện tại là Golang, một ngôn ngữ scripting và khá mới trên thị trường, nên việc tìm nhân sự sẽ gặp ít nhiều khó khăn so với các ngôn ngữ khác như Java, nhất là ở cấp độ Senior. Việc tổ chức code của  một scripting language cũng sẽ không trong sáng như các ngôn ngữ OOP, khi việc xử lý cho business logic ngày càng nhiều lên. - Người Lạ -

Trên đây là một hiểu biết khá sai lầm về Go.

Về cơ bản, Go là một compiled, concurrent, garbage-collected, statically typed language được phát triển tại Google, Open Source và cộng đồng đang phát triển mạnh mẽ, rất nhiều công ty lớn nhỏ trên thế giới đang nghiên cứu và sử dụng nó cho hệ thống của mình. Và tin tôi đi, ít nhiều những thứ bạn đang xài cũng sẽ có chút gì đó liên quan đến Go, bạn đã từng nghe đến Docker, Kubernetes, Prometheus, InfluxDB? Tất cả đều được viết bằng Go.

Go được xây dựng dựa trên 3 tiêu chí **Effecient**, **Scalable**, **Productive**, nó được sinh ra để giải quyết bài toán ở chính Google, bài toán về multicore processor, network, và những hệ thống clusters khổng lồ, với những web programming model né tránh thay vì đối đầu trực diện với vấn đề.

Đầu tiên, build time là một điều cực kỳ quan trọng, một project với hàng triệu dòng code cùng hàng trăm kỹ sư có thể làm việc trên đó, thay đổi mỗi ngày và liên tục, thời gian để clean build nếu mất hàng giờ đó là một vấn đề tồi tệ. Go được thiết kế để quản lý `dependencies` thông minh hơn, `dependency analysis` dễ dàng hơn, và giảm thiểu overheads trong quá trình include, preprocessing headers.

Một vấn đề nữa của Google chính là string processing, Google đọc và phân tích hàng triệu thậm chí hàng tỉ trang web, họ phải xử lý chuỗi rất nhiều, bên trong Go được tích hợp sẵn rất nhiều library với hàng tá hàng hỗ trợ cho việc thao tác với các string. Ngoài ra với việc Garbage collector được tích hợp làm cho việc làm việc với strings dễ dàng và hiệu quả hơn.

Go có một concurrency model rất tuyệt, điều tôi muốn nói đến chính là Goroutine. Có thể nó không đao to búa lớn, không ở cấp độ thần thánh về mặt thiết kế, nhưng ở khía cạnh end-users, nó đơn giản, trực quan và hiệu quả. Vậy còn data races thì sao? Hãy nhìn qua một chút, bạn có thể dùng phương pháp truyền thống là `sync.Mutex`, hoặc hãy xài một tính năng thời thượng đó chính là `channel` ngay hôm nay đi.

### Go là một ngôn ngữ tồi!

Bắt đầu từ những vấn đề thực tế, Go được tạo ra để giải quyết những bài toán đến từ real world, tập trung nhiều vào software engineering hơn là research programming language design, nhiều người sẽ cảm thấy vui vẻ để sử dụng nó như một công cụ mạnh mẽ, nhưng số khác thì cho rằng nó là một `boring language`, không sáng tạo, cú pháp dài dòng, xấu xí.

Bạn sẽ không có những siêu cú pháp ảo diệu để biểu diễn cho thằng ngồi đối diện thấy là "ngạc nhiên chưa?", không có sẵn những syntactic sugar giúp bạn thực hiện những kỹ thuật phức tạp, mọi thứ luôn tuân theo một khuôn mẫu, tuân theo từng cú pháp đơn giản, đôi lúc lại nhàm chán. Thay vào đó bạn phải gần gũi nhiều hơn với nền tảng điện toán, nào là CPU, memory, con trỏ, từng cái buffer, từng bit nhớ...

Tôi vẫn dùng Go, có thể tôi chỉ một lập trình viên tồi, nhưng biết đâu tôi lại là một engineer tốt thì sao? Và vì vẻ đẹp của điện toán nằm ở sự đơn giản.

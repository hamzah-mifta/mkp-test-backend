## System Design Test Answer

Menurut saya, solusi untuk sistem pemilihan tempat duduk adalah dengan menyediakan field `status` pada table `seats` yang menandakan status ketersediaan kursi tersebut.

Pada saat user/customer memilih tempat duduk untuk pemesanan, sistem akan langsung mengubah status tempat duduk tersebut dari `available` menjadi `reserved`.

Tempat duduk yang berstatus `reserved` akan bertahan selama beberapa menit atau jam tergantung dari durasi pembayaran yang disediakan atau bisa juga disesuaikan dengan kebutuhan bisnis.

Status `reserved` tidak akan bisa dipilih oleh user lain saat melakukan pemesanan, sehingga user tidak perlu khawatir tempat duduk yang telah dipilih akan dipesan oleh user lain dan user tidak merasa diburu-buru oleh waktu pembayaran.

Saat user memesan tempat duduk, sistem akan langsung mengubah status ketersediaan menjadi `reserved` di level database. Proses ini juga bisa memanfaatkan cache dengan expiry duration selama waktu yang sebelumnya telah ditentukan untuk mengoptimasi performa sistem. Dengan begitu, user tidak perlu khawatir tempat duduk yang dipilih akan dipesan oleh user lain.

Setelah pembayaran selesai, sistem akan mengubah status tempat duduk menjadi `occupied` yang menandakan kursi tersebut sudah diisi.

Untuk proses restock tiket yang telah terjual, saya berpikir sistem akan mengubah semua status tempat duduk menjadi `available` pada saat film selesai tayang.

Cara lain untuk melakukan restock tiket juga bisa dilakukan secara otomatis oleh sistem dengan memanfaatkan cron dan pubsub. Cron dijalankan per 5 menit untuk mempublish message film mana saja yang sudah selesai tayang. Kemudian consumer akan meng-consume message tersebut dan melakukan batch update status tempat duduk menjadi `available`.

Sedangkan untuk proses refund/pembatalan, tempat duduk yang sebelumnya berstatus `occupied` akan diubah menjadi `available`. Sistem akan memproses pengembalian data kepada user, dengan alasan pembatalan agar customer atau pihak bioskop bisa mengetahui alasan pembatalan dan jumlah refund. Jika pembatalan dilakukan pada saat status `reserved`, status tempat duduk akan langsung dikembalikan menjadi `available` tanpa proses refund.

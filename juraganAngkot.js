// membuat Object Angkot
function Angkot(sopir, trayek, penumpang, kas ) {
    this.sopir = sopir;
    this.trayek = trayek;
    this.penumpang = penumpang;
    this.kas = kas;

    this.penumpangNaik = function(namaPenumpang) {
        this.penumpang.push(namaPenumpang);
        return this.penumpang;
    }

    this.penumpangTurun = function(namaPenumpang, bayar) {
        if(this.penumpang.length == 0) {
            alert('angkot masih kosong')
            return false;
        }

        for (var i = 0; i < this.penumpang.length; i++) {
            if(this.penumpang[i] == namaPenumpang){
                this.penumpang[i] = undefined;
                this.kas += bayar;
                return this.penumpang;
            } else if (this.penumpang[i] == undefined) {
                console.log('Penumpang tidak ada');
                return this.penumpang
            }
        }

    }

}

var angkot1 = new Angkot('Budi', ['Palembang','Plaju'], [],0);
var angkot2 = new Angkot('Michael',['Ilir','Swasembada'],[],0);



function buscarMeusAtivos (capital) {
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var requestOptions = {
        method: 'GET',
        headers: myHeaders,
        redirect: 'follow'
    };

    fetch("http://localhost:8082/", requestOptions)
        .then(response => response.text())
        .then(result => ProcessarUmAUm(result, capital))
        .catch(error => console.log('error', error))
}

function ProcessarUmAUm(todos, capital) {

    const cdbs_ipca = JSON.parse(todos).filter((ativo) => {
        if(ativo.product === "CDB" && ativo.indexers === "Inflação" ) {
            return ativo
        }
    })

    const resposta = cdbs_ipca.map((ativo) => {
        return chamarAPIGo(ativo, capital)
    })

    Promise
        .all(resposta)
        .then((data) => sort(data))
        .catch((error) => console.log(error))
}

function getPorcentagem(string_com_porcentagem) {
    const regex = /\b(?<!\.)(?!0+(?:\.0+)?%)(?:\d|[1-9]\d|100)(?:(?<!100)\,\d+)?%/g

    let percent = regex.exec(string_com_porcentagem)
    let p = percent[0].replace("%", "")
    p = p.replace(",", ".")
    return parseFloat(p)
}

async function chamarAPIGo(ativo, capital) {
    ativo.desc = ativo.fee
    ativo.fee = getPorcentagem(ativo.fee)
    ativo.maturityDate = `${ativo.maturityDate}Z`
    ativo.graceDate = `${ativo.graceDate}Z`

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: JSON.stringify({capital: capital, ativo: ativo})
    };

    return new Promise((resolve,reject)=>{
        fetch("http://127.0.0.1:9002/profit", requestOptions)
        .then(result => result.json())
        .then((json) => {
            result2 = json
            result2.ativo = ativo.nickName
            result2.desc = ativo.desc
            // console.log(result2)
            resolve(result2)
        })
    })
}

function compare( a, b ) {
    if ( a.PercentageProfit < b.PercentageProfit ){
    return 1;
    }
    if ( a.PercentageProfit > b.PercentageProfit ){
    return -1;
    }

    return 0;
}

function sort(todos) {

    ordenados = todos.sort(compare)

    console.log('ordenados', ordenados);
}
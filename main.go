package main

import (
    "log"
    "fmt"
    "encoding/json"
    "math/big"
    "strings"
    "strconv"
    "net/http"
    "./httpUtils"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    StarNFT "./contracts"
)

type Response struct {
    Owner string `json:"owner"`
}

const PORT = 8080

var token StarNFT.StarNFT
var conn *ethclient.Client

func getOwnerOfToken(address string, tokenIdStr string) (result string, err error) {
    token, err := StarNFT.NewStarNFT(common.HexToAddress(address), conn)
    if err != nil {
        return "", fmt.Errorf("Failed to instantiate a Token contract, %v", err)
    }

    // NTF owner
    tokenId := new(big.Int)
    tokenId, ok := tokenId.SetString(tokenIdStr, 10)
    if !ok {
        return "", fmt.Errorf("Invalid tokenId, %v", err)
    }

    owner, err := token.GetOwner(nil, tokenId) 
    if err != nil {
        return "", fmt.Errorf("Failed to retrieve NTF owner, %v", err)
    }

    ownerStr, _ := json.MarshalIndent(owner, "", " ")

    return string(ownerStr), nil
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
    log.Println("Incoming Request:", r.Method)

    switch r.Method {
    case http.MethodGet:
        if r.URL.Path == "/favicon.ico" {
            break
        }

        log.Println(" req: %s", r.URL.Path)
        path := strings.Split(r.URL.Path, "/")
        log.Println("address:" + path[1]);

        if len(path) < 3 || path[1] == "" || path[2] == "" {
            httpUtils.HandleError(&w, 500, "Bad Request", "Contract Address or ID not specified", nil)
        } else {
            owner, err := getOwnerOfToken(path[1], path[2])
            if err != nil {
                httpUtils.HandleError(&w, 500, "Get owner fail:", err.Error(), nil)
            } else {
                resp := Response{ Owner: owner }
                httpUtils.HandleSuccess(&w, resp)
            }
        }

        break
    default:
        httpUtils.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
        break
    }
}

func main() {
    
    // connect to ethereum network
    _conn, err := ethclient.Dial("https://ropsten.infura.io")
    if err != nil {
        log.Fatalf("Connect fail", err)
    }
    conn = _conn

    // launch http server
    log.Println("Attempting to start HTTP Server." + strconv.Itoa(PORT))

    http.HandleFunc("/", HandleRequest)
    err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
    if err != nil {
        log.Panicln("Server failed starting. Error: %s", err)
    }
}
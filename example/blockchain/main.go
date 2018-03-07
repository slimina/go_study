package main

/**
区块链记录心跳
*/
import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/davecgh/go-spew/spew" //在 console 中直接查看 struct 和 slice 这两种数据结构
	"github.com/gorilla/mux"          //Gorilla 的 mux 包非常流行， 我们用它来写 web handler
	"github.com/joho/godotenv"        //godotenv 可以帮助我们读取项目根目录中的 .env 配置文件
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

//数据结构
type Block struct {
	Index     int    //这个块在整个链中的位置
	Timestamp string //块生成时的时间戳
	BPM       int    //每分钟心跳数
	Nonce     string //随机数
	Hash      string //这个块通过 SHA256 算法生成的散列值
	PrevHash  string //前一个块的 SHA256 散列值
}

//区块链
var Blockchain []Block

//计算数据的SHA256 散列值
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//生成block
//必须满足一定条件的新区块才有效 ： 对最新的区块头进行两次SHA256计算，得到的256bit哈希结果，高位48bit必须是0x00000000FFFF，才算挖矿成功。
func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.Nonce = strconv.FormatFloat(rand.Float64(), 'E', -1, 64)
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock, nil
}

// 检查当前块的 Hash 值是否正确
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

//通常来说，更长的链表示它的数据（状态）是更新的，所以我们需要一个函数
//能帮我们将本地的过期的链切换成最新的链
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

//定义路由
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET") //查看块
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")   //添加块
	return muxRouter
}

//初始化web服务
func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

//查看整个链
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		io.WriteString(w, string(bytes))
	}
}

// POST 请求的 payload
type Message struct {
	BPM int
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		newBlockChain := append(Blockchain, newBlock)
		replaceChain(newBlockChain)
		spew.Dump(Blockchain)
	}
	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

//返回response
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	} else {
		w.WriteHeader(code)
		w.Write(response)
	}
}

//主函数
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), 0, "", "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()
	log.Fatal(run())
}

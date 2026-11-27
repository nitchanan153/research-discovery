package main

import (
 "encoding/json"
 "log"
 "net/http"
 "os"
 "time"
)

type Opportunity struct { Disease string `json:"disease"`; Target string `json:"target"`; Modality string `json:"modality"`; Evidence string `json:"evidence"`; Stage string `json:"stage"`; Score int `json:"score"`; Trend string `json:"trend"` }
type Dashboard struct { UpdatedAt string `json:"updatedAt"`; Stats map[string]int `json:"stats"`; Opportunities []Opportunity `json:"opportunities"` }
var data = Dashboard{UpdatedAt: time.Now().UTC().Format(time.RFC3339), Stats: map[string]int{"opportunities":24,"evidence":1284,"technologyTrends":12,"pipeline":8}, Opportunities: []Opportunity{
 {"มะเร็งตับ (HCC)","GPC3","CAR-T Cell Therapy","42 papers · 8 trials","Pre-clinical",92,"+18%"},
 {"โรคอ้วน / เบาหวานชนิดที่ 2","GIPR / GLP-1R","Dual Agonist Peptide","128 papers · 24 trials","Clinical Phase II",86,"+12%"},
 {"มะเร็งปอดชนิด NSCLC","KRAS G12C","Monoclonal Antibody","76 papers · 16 trials","Clinical Phase I",78,"+8%"},
 {"โรคภูมิคุ้มกันทำลายตนเอง","TL1A","Bispecific Antibody","31 papers · 5 trials","Discovery",74,"+24%"},
 {"โรคกล้ามเนื้อเสื่อม (DMD)","Dystrophin","AAV Gene Therapy","54 papers · 11 trials","Clinical Phase III",69,"+5%"},
}}
func main() {
 mux:=http.NewServeMux()
 mux.HandleFunc("GET /api/health",func(w http.ResponseWriter,r *http.Request){writeJSON(w,map[string]string{"status":"ok","service":"bio-rd-api"})})
 mux.HandleFunc("GET /api/v1/dashboard",func(w http.ResponseWriter,r *http.Request){writeJSON(w,data)})
 mux.HandleFunc("GET /api/v1/opportunities",func(w http.ResponseWriter,r *http.Request){q:=r.URL.Query().Get("q"); out:=make([]Opportunity,0); for _,o:=range data.Opportunities { if q==""||containsFold(o.Disease+" "+o.Target+" "+o.Modality+" "+o.Stage,q){out=append(out,o)} }; writeJSON(w,out) })
 addr:=os.Getenv("BIO_RD_ADDR");if addr==""{addr=":8080"}; log.Printf("Bio-RD API listening on %s",addr); log.Fatal(http.ListenAndServe(addr,withCORS(mux)))
}
func writeJSON(w http.ResponseWriter,v any){w.Header().Set("Content-Type","application/json; charset=utf-8");_ = json.NewEncoder(w).Encode(v)}
func withCORS(next http.Handler)http.Handler{return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){w.Header().Set("Access-Control-Allow-Origin","*");w.Header().Set("Access-Control-Allow-Methods","GET, OPTIONS");w.Header().Set("Access-Control-Allow-Headers","Content-Type");if r.Method==http.MethodOptions{w.WriteHeader(http.StatusNoContent);return};next.ServeHTTP(w,r)})}
func containsFold(haystack,needle string)bool{return len(needle)==0||stringContainsFold(haystack,needle)}
func stringContainsFold(haystack,needle string)bool{return len([]rune(needle))==0||indexFold([]rune(haystack),[]rune(needle))>=0}
func indexFold(haystack,needle []rune)int{for i:=0;i+len(needle)<=len(haystack);i++{ok:=true;for j:=range needle{a,b:=haystack[i+j],needle[j];if a>='A'&&a<='Z'{a+=32};if b>='A'&&b<='Z'{b+=32};if a!=b{ok=false;break}};if ok{return i}};return -1}

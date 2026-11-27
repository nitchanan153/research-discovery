# แผนพัฒนาระบบ Bio-RD Research Discovery

> เอกสารนี้แปลงแนวคิดจากภาพ “Bio-RD intelligence & Discovery platform” เป็นแผนพัฒนาระบบ โดยกำหนด Frontend เป็น Next.js และ Backend เป็น Go (Golang) รายละเอียดทางเทคนิคที่ไม่ได้ระบุในภาพถือเป็นข้อเสนอเบื้องต้นและควรยืนยันร่วมกับทีม

## 1. วิสัยทัศน์และเป้าหมาย

สร้างแพลตฟอร์ม Intelligence สำหรับรวบรวม วิเคราะห์ และเชื่อมโยงข้อมูล Bio-pharmaceutical, Biotechnology และข้อมูลที่เกี่ยวข้อง เพื่อค้นหา Research Opportunity และสนับสนุนการตัดสินใจลงทุนวิจัย ตั้งแต่ Disease/Target scouting ไปจนถึง Research Brief, Pre-clinical Roadmap และ Go/Hold/Stop decision

เป้าหมายการใช้งาน:
- ลดเวลาค้นหาและรวบรวม Literature, Patent, Clinical Trial และข้อมูลจากหลายแหล่ง
- เชื่อม Disease → Gene/Protein/Target → Biological Drug/Modality → Evidence
- ระบุ Research Gap, Technology Gap และ Unmet Medical Need
- จัดอันดับโอกาสวิจัยโดยอิงหลักฐาน ความเป็นไปได้ และยุทธศาสตร์องค์กร
- สร้างรายงานและแผนวิจัยที่ตรวจสอบย้อนกลับไปยังแหล่งข้อมูลได้

## 2. ผู้ใช้และผู้มีส่วนเกี่ยวข้อง

- Technology Intelligence (เจ้าของโครงการ)
- Biologic R&D และ Biotechnologist
- Data Science และ Bioinformatics
- Business Development
- QA, RA และ IP
- ผู้เชี่ยวชาญ/ที่ปรึกษาทางวิทยาศาสตร์และ Research Partner
- ผู้ดูแลระบบและผู้ดูแลข้อมูล

## 3. ขอบเขตข้อมูลนำเข้า

1. **Disease/Therapeutic Area:** Oncology, Immunology, Metabolic Disease, Rare Disease และสาขาที่องค์กรกำหนด
2. **Biological Drug และ Modality:** Monoclonal Antibody, Recombinant Protein, Peptide, Cell Therapy, Gene Therapy, RNA Therapeutics และ Novel Biologics
3. **Scientific Literature/Evidence:** Research Paper, Review, Pre-clinical Study, Mechanism of Action และ Experimental Evidence
4. **Patent/IP:** Patent Landscape, เจ้าของเทคโนโลยี, อายุสิทธิบัตร และการคัดกรอง Freedom-to-Operate เบื้องต้น
5. **Clinical/Pre-clinical Development:** Drug Pipeline, สถานะ Clinical Trial, Development Stage และหลักฐานก่อนคลินิก
6. **ยุทธศาสตร์และความสามารถองค์กร:** Research Priority, Technology Gap, Laboratory Capability, Budget และ Strategic Direction

ข้อมูลจากแหล่งภายนอกต้องเก็บชื่อแหล่งที่มา URL/รหัสระเบียน วันที่ดึงข้อมูล เวอร์ชัน และเงื่อนไขการใช้งาน เพื่อให้ตรวจสอบย้อนหลังได้

## 4. ความสามารถหลักของระบบ

### 4.1 Research Intelligence Dashboard
- สรุป Disease Area, Biological Drug, Technology Trend และ Research Pipeline
- ตัวกรองตามโรค เป้าหมาย โมดาลิตี ระยะพัฒนา บริษัท/สถาบัน และช่วงเวลา
- แสดงตัวชี้วัดจำนวนหลักฐาน แนวโน้ม และสถานะข้อมูลล่าสุด

### 4.2 Disease & Target Opportunity Map
- แสดงความสัมพันธ์ระหว่าง Disease, Gene, Protein, Target, Mechanism of Action, Drug Candidate และ Evidence
- ให้ผู้ใช้เปิดดูหลักฐานและที่มาของแต่ละความสัมพันธ์
- แสดงช่องว่างข้อมูลและความเชื่อมั่นของความสัมพันธ์อย่างชัดเจน

### 4.3 Technology Landscape Report
- เปรียบเทียบ Biological Modality, Development Stage และ Competitive Landscape
- สรุปแนวโน้มเทคโนโลยีและผู้พัฒนา/เจ้าของทรัพย์สินทางปัญญา
- ส่งออกรายงานเพื่อใช้ประกอบการประชุมและการประเมิน

### 4.4 Opportunity Scorecard และการจัดลำดับ
- จัดอันดับ Research Opportunity ตาม Scientific Evidence, Novelty, Feasibility และ Strategic Fit
- แสดงคะแนนรายเกณฑ์ น้ำหนักที่ใช้ เหตุผล หลักฐาน และข้อมูลที่ยังขาด
- รองรับการปรับเกณฑ์/น้ำหนักโดยผู้มีสิทธิ์ และบันทึกประวัติการเปลี่ยนแปลง

### 4.5 AI Research Assistant
- ค้นหาและสรุป Paper, Patent และ Clinical Trial ด้วยคำถามภาษาธรรมชาติ
- สกัดข้อมูลสำคัญและเชื่อมโยง Disease, Gene, Protein, Target และ Biological Drug
- ทุกข้อสรุปต้องแนบ citation ที่เปิดกลับไปยังเอกสารหรือส่วนข้อความต้นทางได้
- แสดงสถานะว่าเป็นข้อความจากแหล่งข้อมูล ข้อสรุปจากโมเดล หรือการประเมินของผู้เชี่ยวชาญ
- ผู้ใช้สามารถตรวจแก้และให้คะแนนคำตอบเพื่อปรับปรุงคุณภาพ

### 4.6 Research Brief และ Pre-clinical Roadmap
- สร้าง Research Gap, Hypothesis, Research Question, Experimental Objective และ Expected Outcome
- ช่วยประเมิน Computational หรือ Experimental Model, Equipment, Reagent และ Research Partner ที่ต้องใช้
- วางแผน In-vitro, Ex-vivo หรือ In-vivo ตามความเหมาะสม พร้อม Endpoint และ Evidence ที่ต้องการ
- สรุปเป็น Research Brief, Feasibility Assessment และ Pre-clinical Research Proposal
- รองรับการพิจารณา Go / Hold / Stop พร้อมเหตุผลและผู้อนุมัติ

## 5. Workflow การทำงาน

1. **กำหนดโจทย์:** เลือก Disease/Therapeutic Area, Modality และคำถามวิจัยเบื้องต้น
2. **รวบรวมข้อมูล:** ดึงข้อมูลจากแหล่ง Literature, Patent, Clinical Trial และ Biological Knowledge Base
3. **ทำความสะอาดและเชื่อมโยง:** จัดรูปแบบ entity, แก้รายการซ้ำ, ผูก identifier และบันทึก provenance
4. **Technology & Disease Scouting:** วิเคราะห์เทคโนโลยีเกิดใหม่ Disease Opportunity, Unmet Need และ Research/Technology Gap
5. **สร้าง Evidence Base:** สกัดข้อมูลจากเอกสาร จัดระดับ Research Stage/Evidence Level และเก็บ citation
6. **วิเคราะห์และจัดลำดับ:** ประเมิน Disease Target และ Biological Drug Candidate จัดทำ Opportunity Scorecard
7. **ผู้เชี่ยวชาญทบทวน:** ตรวจความถูกต้องของหลักฐาน ความสอดคล้องกับยุทธศาสตร์ และ IP Risk เบื้องต้น
8. **พัฒนา Research Plan:** สร้าง Hypothesis, Research Question, Experimental Objective และ Feasibility Assessment
9. **วางแผนก่อนคลินิก:** เตรียม Experimental Validation Roadmap และ Pre-clinical Research Proposal
10. **ตัดสินใจ:** Go เมื่อมีหลักฐานและความเป็นไปได้เพียงพอ, Hold เมื่อจำเป็นต้องเติมข้อมูล/ความพร้อม, Stop เมื่อความเสี่ยงหรือข้อจำกัดสูงกว่าคุณค่าที่คาดหวัง
11. **ติดตามและปรับปรุง:** อัปเดตข้อมูลและ pipeline ตามรอบ พร้อมเก็บประวัติการตัดสินใจ

## 6. แหล่งข้อมูลและเครื่องมือที่เสนอ

| กลุ่ม | ตัวอย่างจากแนวคิดต้นทาง | การใช้งาน |
|---|---|---|
| Literature | PubMed, Europe PMC, Semantic Scholar, OpenAlex | Paper, research evidence และแนวโน้มวิทยาศาสตร์ |
| Patent/Clinical | Google Patents, WIPO PATENTSCOPE, ClinicalTrials.gov | Patent landscape, IP และ clinical/pre-clinical pipeline |
| Biological knowledge | UniProt, NCBI Gene, Open Targets, KEGG, STRING | เชื่อม Gene, Protein, Pathway, Target และ Disease |
| AI | LLM + RAG + scientific document extraction | ค้นหา สรุป สกัดข้อมูล และตอบพร้อม citation |
| Analytics | SQL, Python/R หรือ dashboard web-based | Opportunity Score, Research Map และ Intelligence Dashboard |

ก่อนเชื่อมต่อแต่ละแหล่ง ต้องตรวจสอบ API, license, rate limit, ความถี่อัปเดต และสิทธิ์การนำข้อมูลมาใช้จริง แหล่งข้อมูลในตารางเป็นตัวอย่างจากภาพ ไม่ใช่การยืนยันว่ามีสิทธิ์หรือ API พร้อมใช้งาน

## 7. สถาปัตยกรรมที่เสนอ

```text
ผู้ใช้
  │
  ▼
Next.js Web Application
  ├─ Dashboard / Search / Opportunity Map
  ├─ Evidence Viewer / Reports / Workflow
  └─ เรียก API ผ่าน HTTPS
       │
       ▼
Go Backend (REST API)
  ├─ Identity & Access / Project / Search
  ├─ Research domain services / Scoring / Workflow
  ├─ Ingestion connectors / Background jobs
  └─ AI orchestration และการจัดการ citation
       │
       ├─ PostgreSQL: ข้อมูลธุรกิจและ normalized entities
       ├─ Object Storage: PDF, ไฟล์แนบ และรายงาน
       ├─ Search index: full-text/semantic retrieval (เลือกใช้ตามผลประเมิน)
       ├─ Queue/worker: ingestion, extraction และงานใช้เวลานาน
       └─ External scientific data sources / LLM provider
```

ข้อเสนอทางเทคนิคเบื้องต้น:
- **Frontend:** Next.js, TypeScript, component library ที่ทีมเลือก, server/client rendering ตามหน้าจอ, responsive layout
- **Backend:** Go แยกเป็น domain modules ภายในบริการเดียวในระยะแรก (modular monolith) และเปิด REST API พร้อม OpenAPI contract
- **ฐานข้อมูล:** PostgreSQL เป็นแหล่งข้อมูลหลัก; ใช้ full-text search ก่อน และเพิ่ม vector/search engine เมื่อมี use case และข้อมูลประเมินชัดเจน
- **ไฟล์:** Object storage พร้อม metadata, checksum, access control และนโยบาย retention
- **งานเบื้องหลัง:** worker และ queue สำหรับดึงข้อมูล/ประมวลผลเอกสารที่ไม่ควรผูกกับ request ผู้ใช้
- **AI/RAG:** ให้บริการผ่าน adapter แยกจาก business logic เพื่อเปลี่ยน provider/model ได้; บันทึก prompt/model version และแหล่งอ้างอิงตามนโยบายองค์กร
- **การ deploy:** Containerized environment, CI/CD, secrets management, monitoring, centralized logging และ backup/restore

## 8. โมดูล Backend และข้อมูลหลัก

โมดูล Go ที่แนะนำ:
1. `identity`: ผู้ใช้ บทบาท และสิทธิ์
2. `workspace`: โครงการ research, สมาชิก, tag และสถานะ
3. `catalog`: Disease, Gene, Protein, Target, Modality, Drug Candidate และ Organization
4. `sources`: แหล่งข้อมูล, ingestion, sync status และ provenance
5. `evidence`: เอกสาร ข้อความที่สกัด citation และระดับหลักฐาน
6. `discovery`: ความสัมพันธ์, opportunity, research gap และ landscape
7. `scoring`: เกณฑ์ น้ำหนัก คะแนน และคำอธิบายผล
8. `assistant`: retrieval, summarization และคำตอบพร้อมแหล่งอ้างอิง
9. `planning`: brief, hypothesis, roadmap, feasibility และ Go/Hold/Stop
10. `reporting`: dashboard query, export และรายงาน
11. `audit`: ประวัติการแก้ไข การทบทวน และการตัดสินใจ

Entity หลักควรมี stable ID และรองรับ alias/synonym รวมถึง external ID ตามแหล่งข้อมูล ตัวอย่าง entity: `Disease`, `Gene`, `Protein`, `Target`, `Modality`, `DrugCandidate`, `Evidence`, `Patent`, `ClinicalTrial`, `Organization`, `Opportunity`, `ResearchBrief`, `Decision` และ `SourceRecord`

## 9. หน้าจอ Next.js ที่ควรพัฒนา

- เข้าสู่ระบบและเลือก Research Workspace
- Overview Dashboard
- Search & AI Research Assistant
- Disease/Target Opportunity Map
- Evidence Explorer พร้อมตัวกรองและ citation
- Technology/Drug Landscape
- Opportunity Scorecard และหน้ารายละเอียดคะแนน
- Research Brief/Hypothesis Editor
- Pre-clinical Roadmap และ Go/Hold/Stop review
- Data Source/Ingestion Monitor
- Admin: ผู้ใช้ สิทธิ์ เกณฑ์คะแนน และการตั้งค่า

## 10. แผนพัฒนาเป็นระยะ

| ระยะ | งานหลัก | ผลส่งมอบ |
|---|---|---|
| 0. Discovery & Governance | เก็บ user journey, use case, data source, licensing, security และนิยามคะแนน | Product scope, data inventory, architecture decision และ acceptance criteria |
| 1. Platform Foundation | ตั้ง repo, CI/CD, auth/RBAC, workspace, audit, database และ UI shell | ระบบพื้นฐานพร้อมใช้งานภายใน |
| 2. Data Foundation | เชื่อมแหล่งข้อมูลชุดแรก, ingestion worker, normalization, deduplication และ provenance | Evidence database ที่อัปเดตและตรวจสอบย้อนกลับได้ |
| 3. Discovery MVP | Search, entity pages, dashboard พื้นฐาน, evidence explorer และ opportunity map | ผู้ใช้ค้น Disease/Target และเปิดดูหลักฐานได้ |
| 4. Intelligence | Document extraction, AI assistant พร้อม citation, landscape และ scorecard | สรุปและจัดอันดับ opportunity ที่มีคำอธิบายประกอบ |
| 5. Research Planning | Brief, hypothesis, feasibility, roadmap, review และ Go/Hold/Stop | กระบวนการเปลี่ยน intelligence เป็นแผนวิจัย |
| 6. Production Readiness | ประเมินความปลอดภัย ประสิทธิภาพ คุณภาพข้อมูล monitoring, backup และการอบรม | แผนเปิดใช้จริงและคู่มือปฏิบัติงาน |

ควรเริ่ม MVP ด้วย Disease Area จำกัด 1–2 สาขา และแหล่งข้อมูลที่มีสิทธิ์ใช้งานชัดเจน เพื่อพิสูจน์ workflow ก่อนขยาย coverage

## 11. ข้อกำหนดด้านคุณภาพและการควบคุม

- ทุกข้อสรุปจาก AI ต้องมีแหล่งอ้างอิงและเปิดดูหลักฐานต้นทางได้
- แสดงวันที่อัปเดต ความครอบคลุม และข้อจำกัดของข้อมูล
- แยก evidence, model-generated summary และ expert assessment
- มี human review สำหรับการจัดอันดับและการตัดสินใจลงทุน/วิจัย
- ใช้ role-based access control, audit log, encryption ระหว่างส่ง/จัดเก็บ และ secrets management
- ประเมินการจัดการข้อมูลอ่อนไหว ข้อมูลส่วนบุคคล และข้อกำหนด IP ตามนโยบายองค์กร
- วางแผนสำรอง/กู้คืน กำหนด retention และจัดการการลบ/อัปเดตข้อมูลจากต้นทาง
- รองรับการค้นหา การกรอง และการเปิดเอกสารโดยไม่ทำให้ผู้ใช้เข้าใจคะแนนเป็นข้อสรุปทางการแพทย์หรือข้อกฎหมาย

## 12. ตัวชี้วัดความสำเร็จ

- เวลามัธยฐานในการจัดทำ landscape/research brief ลดลงจาก baseline ที่องค์กรวัดได้
- สัดส่วนระเบียนที่มี source, identifier และวันที่อัปเดตครบถ้วน
- ความแม่นยำของการสกัด entity/citation จากชุดข้อมูลตัวอย่างที่ผู้เชี่ยวชาญตรวจ
- สัดส่วนคำตอบ AI ที่มี citation ใช้ตรวจสอบได้ และอัตราคำตอบที่ผู้เชี่ยวชาญยอมรับ
- เวลาที่ใช้จากการตั้งโจทย์จนได้ shortlist และการตัดสินใจ Go/Hold/Stop
- การใช้งานซ้ำของทีมวิจัยและสัดส่วน opportunity ที่มีผู้รับผิดชอบ/แผนขั้นถัดไป

กำหนดค่าเป้าหมายหลังเก็บ baseline และทำ pilot โดยไม่สมมติค่าความแม่นยำหรือผลประหยัดล่วงหน้า

## 13. ความเสี่ยงและแนวทางลดความเสี่ยง

| ความเสี่ยง | แนวทาง |
|---|---|
| สิทธิ์ใช้ข้อมูลหรือ API เปลี่ยนแปลง | ตรวจ license/เงื่อนไขก่อนเชื่อมต่อ จัดทำ source registry และรองรับการปิด connector |
| ข้อมูลซ้ำ ชื่อไม่ตรง หรือ entity เชื่อมผิด | ใช้ external identifier, synonym mapping, confidence และ human review |
| AI สรุปคลาดเคลื่อนหรือสร้าง citation ที่ไม่มีจริง | จำกัดคำตอบให้ยึด retrieved evidence, ตรวจ citation และเปิดเอกสารต้นทาง |
| คะแนน Opportunity ดูเป็นวัตถุวิสัยเกินจริง | แสดงน้ำหนัก เหตุผล ความไม่แน่นอน และให้ผู้เชี่ยวชาญอนุมัติ |
| งานนำเข้า/ประมวลผลใช้เวลานาน | ใช้ asynchronous jobs, retry, idempotency และ ingestion monitoring |
| IP หรือข้อมูลภายในรั่วไหล | แยกสิทธิ์ workspace, ตรวจผู้ให้บริการ AI, จำกัดข้อมูลส่งออก และบันทึก audit |
| ขอบเขตใหญ่เกิน MVP | เริ่มจาก use case/โรคจำกัดและทบทวนผล pilot ก่อนขยาย |

## 14. สิ่งที่ต้องตัดสินใจในช่วงเริ่มโครงการ

- Disease Area แรกและกลุ่มผู้ใช้หลักของ MVP
- แหล่งข้อมูลที่องค์กรมีสิทธิ์ใช้งานและความถี่ sync
- ระบบยืนยันตัวตนและการเชื่อมต่อกับระบบองค์กร
- เกณฑ์/น้ำหนัก Opportunity Score และผู้มีอำนาจอนุมัติ
- ผู้ให้บริการ LLM, ข้อกำหนด data residency และข้อมูลที่ห้ามส่งออก
- Search/indexing, cloud/on-premise และข้อกำหนดความปลอดภัย
- รูปแบบรายงานและการเชื่อมต่อ BI/ระบบจัดการโครงการที่มีอยู่

## 15. เกณฑ์ยอมรับ MVP

- ผู้ใช้ที่ได้รับสิทธิ์สร้าง workspace และกำหนด Disease/Therapeutic Area ได้
- ระบบนำเข้าข้อมูลจากแหล่งที่ตกลงกัน พร้อม source, identifier และวันที่อัปเดต
- ค้นหา Disease/Target/Drug Candidate และเปิดหลักฐานที่เกี่ยวข้องได้
- แสดงความสัมพันธ์และ Opportunity shortlist พร้อมหลักฐาน/เหตุผลประกอบ
- AI ตอบคำถามจากชุดข้อมูลที่ค้นพบและแสดง citation ที่ตรวจสอบได้
- สร้าง Research Brief ฉบับร่างและบันทึกการทบทวน/การตัดสินใจได้
- มี audit trail ขั้นพื้นฐานและผู้ดูแลตรวจสถานะ ingestion ได้

---

**หมายเหตุ:** ภาพต้นทางเป็น workflow และแนวคิดระดับธุรกิจ เอกสารนี้จึงเป็นแผนตั้งต้น ไม่ใช่ข้อกำหนด API, schema, vendor หรือ timeline ที่ได้รับอนุมัติแล้ว ควรยืนยันรายละเอียดในระยะ Discovery & Governance ก่อนเริ่ม implementation

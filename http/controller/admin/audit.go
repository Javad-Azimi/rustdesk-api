package admin

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lejianwen/rustdesk-api/v2/global"
	"github.com/lejianwen/rustdesk-api/v2/http/request/admin"
	"github.com/lejianwen/rustdesk-api/v2/http/response"
	"github.com/lejianwen/rustdesk-api/v2/model"
	"github.com/lejianwen/rustdesk-api/v2/service"
	"gorm.io/gorm"
)

type Audit struct {
}

// ConnList 列表
// @Tags 链接日志
// @Summary 链接日志列表
// @Description 链接日志列表
// @Accept  json
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "页大小"
// @Param peer_id query int false "目标设备"
// @Param from_peer query int false "来源设备"
// @Success 200 {object} response.Response{data=model.AuditConnList}
// @Failure 500 {object} response.Response
// @Router /admin/audit_conn/list [get]
// @Security token
func (a *Audit) ConnList(c *gin.Context) {
	query := &admin.AuditQuery{}
	if err := c.ShouldBindQuery(query); err != nil {
		response.Fail(c, 101, response.TranslateMsg(c, "ParamsError")+err.Error())
		return
	}
	res := service.AllService.AuditService.AuditConnList(query.Page, query.PageSize, func(tx *gorm.DB) {
		if query.PeerId != "" {
			tx.Where("peer_id like ?", "%"+query.PeerId+"%")
		}
		if query.FromPeer != "" {
			tx.Where("from_peer like ?", "%"+query.FromPeer+"%")
		}
		tx.Order("id desc")
	})
	response.Success(c, res)
}

// ConnDelete 删除
// @Tags 链接日志
// @Summary 链接日志删除
// @Description 链接日志删除
// @Accept  json
// @Produce  json
// @Param body body model.AuditConn true "链接日志信息"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/audit_conn/delete [post]
// @Security token
func (a *Audit) ConnDelete(c *gin.Context) {
	f := &model.AuditConn{}
	if err := c.ShouldBindJSON(f); err != nil {
		response.Fail(c, 101, response.TranslateMsg(c, "ParamsError")+err.Error())
		return
	}
	id := f.Id
	errList := global.Validator.ValidVar(c, id, "required,gt=0")
	if len(errList) > 0 {
		response.Fail(c, 101, errList[0])
		return
	}
	l := service.AllService.AuditService.ConnInfoById(f.Id)
	if l.Id > 0 {
		err := service.AllService.AuditService.DeleteAuditConn(l)
		if err == nil {
			response.Success(c, nil)
			return
		}
		response.Fail(c, 101, err.Error())
		return
	}
	response.Fail(c, 101, response.TranslateMsg(c, "ItemNotFound"))
}

// BatchConnDelete 删除
// @Tags 链接日志
// @Summary 链接日志批量删除
// @Description 链接日志批量删除
// @Accept  json
// @Produce  json
// @Param body body admin.AuditConnLogIds true "链接日志"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/audit_conn/batchDelete [post]
// @Security token
func (a *Audit) BatchConnDelete(c *gin.Context) {
	f := &admin.AuditConnLogIds{}
	if err := c.ShouldBindJSON(f); err != nil {
		response.Fail(c, 101, response.TranslateMsg(c, "ParamsError")+err.Error())
		return
	}
	if len(f.Ids) == 0 {
		response.Fail(c, 101, response.TranslateMsg(c, "ParamsError"))
		return
	}

	err := service.AllService.AuditService.BatchDeleteAuditConn(f.Ids)
	if err == nil {
		response.Success(c, nil)
		return
	}
	response.Fail(c, 101, err.Error())
	return
}

// FileList 列表
// @Tags 文件日志
// @Summary 文件日志列表
// @Description 文件日志列表
// @Accept  json
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "页大小"
// @Param peer_id query int false "目标设备"
// @Param from_peer query int false "来源设备"
// @Success 200 {object} response.Response{data=model.AuditFileList}
// @Failure 500 {object} response.Response
// @Router /admin/audit_file/list [get]
// @Security token
func (a *Audit) FileList(c *gin.Context) {
	query := &admin.AuditQuery{}
	if err := c.ShouldBindQuery(query); err != nil {
		response.Fail(c, 101, response.TranslateMsg(c, "ParamsError")+err.Error())
		return
	}
	res := service.AllService.AuditService.AuditFileList(query.Page, query.PageSize, func(tx *gorm.DB) {
		if query.PeerId != "" {
			tx.Where("peer_id like ?", "%"+query.PeerId+"%")
		}
		if query.FromPeer != "" {
			tx.Where("from_peer like ?", "%"+query.FromPeer+"%")
		}
		tx.Order("id desc")
	})
	response.Success(c, res)
}

// FileDelete 删除
// @Tags 文件日志
// @Summary 文件日志删除
// @Description 文件日志删除
// @Accept  json
// @Produce  json
// @Param body body model.AuditFile true "文件日志信息"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/audit_file/delete [post]
// @Security token
func (a *Audit) FileDelete(c *gin.Context) {
	f := &model.AuditFile{}
	if err := c.ShouldBindJSON(f); err != nil {
		response.Fail(c, 101, response.TranslateMsg(c, "ParamsError")+err.Error())
		return
	}
	id := f.Id
	errList := global.Validator.ValidVar(c, id, "required,gt=0")
	if len(errList) > 0 {
		response.Fail(c, 101, errList[0])
		return
	}
	l := service.AllService.AuditService.FileInfoById(f.Id)
	if l.Id > 0 {
		err := service.AllService.AuditService.DeleteAuditFile(l)
		if err == nil {
			response.Success(c, nil)
			return
		}
		response.Fail(c, 101, err.Error())
		return
	}
	response.Fail(c, 101, response.TranslateMsg(c, "ItemNotFound"))
}

// BatchFileDelete 删除
// @Tags 文件日志
// @Summary 文件日志批量删除
// @Description 文件日志批量删除
// @Accept  json
// @Produce  json
// @Param body body admin.AuditFileLogIds true "文件日志"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/audit_file/batchDelete [post]
// @Security token
func (a *Audit) BatchFileDelete(c *gin.Context) {
	f := &admin.AuditFileLogIds{}
	if err := c.ShouldBindJSON(f); err != nil {
		response.Fail(c, 101, response.TranslateMsg(c, "ParamsError")+err.Error())
		return
	}
	if len(f.Ids) == 0 {
		response.Fail(c, 101, response.TranslateMsg(c, "ParamsError"))
		return
	}

	err := service.AllService.AuditService.BatchDeleteAuditFile(f.Ids)
	if err == nil {
		response.Success(c, nil)
		return
	}
	response.Fail(c, 101, err.Error())
	return
}

type auditReportConnItem struct {
	PeerId          string `json:"peer_id"`
	ConnId          int64  `json:"conn_id"`
	FromPeer        string `json:"from_peer"`
	FromName        string `json:"from_name"`
	StaffUsername   string `json:"staff_username"`
	Ip              string `json:"ip"`
	StartTime       int64  `json:"start_time"`
	CloseTime       int64  `json:"close_time"`
	DurationSeconds int64  `json:"duration_seconds"`
}

type auditReportFileItem struct {
	PeerId        string `json:"peer_id"`
	FromPeer      string `json:"from_peer"`
	FromName      string `json:"from_name"`
	StaffUsername string `json:"staff_username"`
	Direction     int    `json:"direction"`
	FileName      string `json:"file_name"`
	FileSizeBytes int64  `json:"file_size_bytes"`
	CreatedAt     int64  `json:"created_at"`
}

type auditReportPeerSummary struct {
	PeerId               string `json:"peer_id"`
	ConnectionCount      int    `json:"connection_count"`
	TotalDurationSeconds int64  `json:"total_duration_seconds"`
	FileCount            int    `json:"file_count"`
	TotalFileBytes       int64  `json:"total_file_bytes"`
}

type auditCustomerReportResponse struct {
	FromTime             int64                     `json:"from_time"`
	ToTime               int64                     `json:"to_time"`
	TotalConnections     int                       `json:"total_connections"`
	TotalDurationSeconds int64                     `json:"total_duration_seconds"`
	TotalFileCount       int                       `json:"total_file_count"`
	TotalFileBytes       int64                     `json:"total_file_bytes"`
	ByPeer               []*auditReportPeerSummary `json:"by_peer"`
	Connections          []*auditReportConnItem    `json:"connections"`
	Files                []*auditReportFileItem    `json:"files"`
}

type auditFileInfoJson struct {
	Files [][]interface{} `json:"files"`
	Name  string          `json:"name"`
	Ip    string          `json:"ip"`
	Num   int             `json:"num"`
}

func (a *Audit) CustomerReport(c *gin.Context) {
	rawPeerIds := strings.TrimSpace(c.Query("peer_ids"))
	now := time.Now()
	fromUnix := now.AddDate(0, -3, 0).Unix()
	toUnix := now.Unix()

	if v := strings.TrimSpace(c.Query("from_time")); v != "" {
		if parsed, err := strconv.ParseInt(v, 10, 64); err == nil && parsed > 0 {
			fromUnix = parsed
		}
	}
	if v := strings.TrimSpace(c.Query("to_time")); v != "" {
		if parsed, err := strconv.ParseInt(v, 10, 64); err == nil && parsed > 0 {
			toUnix = parsed
		}
	}

	peerIds := make([]string, 0)
	peerIdSet := make(map[string]bool)
	for _, p := range strings.Split(rawPeerIds, ",") {
		p = strings.TrimSpace(p)
		if p != "" && !peerIdSet[p] {
			peerIdSet[p] = true
			peerIds = append(peerIds, p)
		}
	}

	resp := &auditCustomerReportResponse{
		FromTime:    fromUnix,
		ToTime:      toUnix,
		ByPeer:      make([]*auditReportPeerSummary, 0),
		Connections: make([]*auditReportConnItem, 0),
		Files:       make([]*auditReportFileItem, 0),
	}

	if len(peerIds) == 0 {
		response.Success(c, resp)
		return
	}

	fromDt := time.Unix(fromUnix, 0)
	toDt := time.Unix(toUnix, 0)

	var rawConns []*model.AuditConn
	global.DB.Where("peer_id IN ? AND created_at >= ? AND created_at <= ?", peerIds, fromDt, toDt).
		Order("id asc").
		Find(&rawConns)

	var rawFiles []*model.AuditFile
	global.DB.Where("peer_id IN ? AND created_at >= ? AND created_at <= ?", peerIds, fromDt, toDt).
		Order("id asc").
		Find(&rawFiles)

	// Collect all from_peer IDs to resolve staff login usernames
	fromPeerSet := make(map[string]bool)
	for _, rc := range rawConns {
		if rc.FromPeer != "" && rc.FromPeer != "web" {
			fromPeerSet[rc.FromPeer] = true
		}
	}
	for _, rf := range rawFiles {
		if rf.FromPeer != "" && rf.FromPeer != "web" {
			fromPeerSet[rf.FromPeer] = true
		}
	}

	peerToUsername := make(map[string]string)
	if len(fromPeerSet) > 0 {
		fromPeers := make([]string, 0, len(fromPeerSet))
		for fp := range fromPeerSet {
			fromPeers = append(fromPeers, fp)
		}
		type peerUserRow struct {
			PeerId   string
			Username string
		}
		var rows []peerUserRow
		global.DB.Table("peers").
			Select("peers.id as peer_id, users.username as username").
			Joins("left join users on users.id = peers.user_id").
			Where("peers.id IN ?", fromPeers).
			Scan(&rows)
		for _, r := range rows {
			if r.Username != "" {
				peerToUsername[r.PeerId] = r.Username
			}
		}
	}

	// Initialize per-peer summary map
	summaryMap := make(map[string]*auditReportPeerSummary)
	for _, pid := range peerIds {
		summaryMap[pid] = &auditReportPeerSummary{PeerId: pid}
	}

	// Deduplicate connections by (peer_id, conn_id)
	connMap := make(map[string]*auditReportConnItem)
	for _, rc := range rawConns {
		key := fmt.Sprintf("%s:%d", rc.PeerId, rc.ConnId)
		startUnix := time.Time(rc.CreatedAt).Unix()
		existing, ok := connMap[key]
		if !ok {
			staffUser := peerToUsername[rc.FromPeer]
			if staffUser == "" && rc.FromName != "" && rc.FromName != "web" {
				staffUser = rc.FromName
			}
			item := &auditReportConnItem{
				PeerId:        rc.PeerId,
				ConnId:        rc.ConnId,
				FromPeer:      rc.FromPeer,
				FromName:      rc.FromName,
				StaffUsername: staffUser,
				Ip:            rc.Ip,
				StartTime:     startUnix,
				CloseTime:     rc.CloseTime,
			}
			connMap[key] = item
		} else {
			if startUnix > 0 && (existing.StartTime == 0 || startUnix < existing.StartTime) {
				existing.StartTime = startUnix
			}
			if rc.CloseTime > existing.CloseTime {
				existing.CloseTime = rc.CloseTime
			}
			if existing.FromPeer == "" && rc.FromPeer != "" {
				existing.FromPeer = rc.FromPeer
			}
			if (existing.FromName == "" || existing.FromName == "web") && rc.FromName != "" {
				existing.FromName = rc.FromName
			}
			if existing.StaffUsername == "" {
				if u := peerToUsername[rc.FromPeer]; u != "" {
					existing.StaffUsername = u
				} else if rc.FromName != "" && rc.FromName != "web" {
					existing.StaffUsername = rc.FromName
				}
			}
		}
	}

	for _, item := range connMap {
		if item.CloseTime > item.StartTime && item.StartTime > 0 {
			item.DurationSeconds = item.CloseTime - item.StartTime
		}
		resp.Connections = append(resp.Connections, item)
		resp.TotalConnections++
		resp.TotalDurationSeconds += item.DurationSeconds
		if s, ok := summaryMap[item.PeerId]; ok {
			s.ConnectionCount++
			s.TotalDurationSeconds += item.DurationSeconds
		}
	}

	sort.Slice(resp.Connections, func(i, j int) bool {
		return resp.Connections[i].StartTime > resp.Connections[j].StartTime
	})

	// Deduplicate audit_files retries (same peer_id + from_peer + info within 60s)
	lastFileSeen := make(map[string]int64)
	for _, rf := range rawFiles {
		ts := time.Time(rf.CreatedAt).Unix()
		dedupKey := rf.PeerId + "|" + rf.FromPeer + "|" + rf.Info
		if prevTs, seen := lastFileSeen[dedupKey]; seen && (ts-prevTs) <= 60 {
			continue
		}
		lastFileSeen[dedupKey] = ts

		staffUser := peerToUsername[rf.FromPeer]
		if staffUser == "" && rf.FromName != "" && rf.FromName != "web" {
			staffUser = rf.FromName
		}

		var parsed auditFileInfoJson
		if err := json.Unmarshal([]byte(rf.Info), &parsed); err == nil && len(parsed.Files) > 0 {
			for _, fEntry := range parsed.Files {
				if len(fEntry) < 2 {
					continue
				}
				fname, _ := fEntry[0].(string)
				var fsize int64
				switch v := fEntry[1].(type) {
				case float64:
					fsize = int64(v)
				case int64:
					fsize = v
				}
				fileItem := &auditReportFileItem{
					PeerId:        rf.PeerId,
					FromPeer:      rf.FromPeer,
					FromName:      rf.FromName,
					StaffUsername: staffUser,
					Direction:     rf.Type,
					FileName:      fname,
					FileSizeBytes: fsize,
					CreatedAt:     ts,
				}
				resp.Files = append(resp.Files, fileItem)
				resp.TotalFileCount++
				resp.TotalFileBytes += fsize
				if s, ok := summaryMap[rf.PeerId]; ok {
					s.FileCount++
					s.TotalFileBytes += fsize
				}
			}
		} else if rf.Path != "" {
			fileItem := &auditReportFileItem{
				PeerId:        rf.PeerId,
				FromPeer:      rf.FromPeer,
				FromName:      rf.FromName,
				StaffUsername: staffUser,
				Direction:     rf.Type,
				FileName:      rf.Path,
				FileSizeBytes: 0,
				CreatedAt:     ts,
			}
			resp.Files = append(resp.Files, fileItem)
			resp.TotalFileCount++
			if s, ok := summaryMap[rf.PeerId]; ok {
				s.FileCount++
			}
		}
	}

	sort.Slice(resp.Files, func(i, j int) bool {
		return resp.Files[i].CreatedAt > resp.Files[j].CreatedAt
	})

	for _, pid := range peerIds {
		if s, ok := summaryMap[pid]; ok {
			resp.ByPeer = append(resp.ByPeer, s)
		}
	}

	response.Success(c, resp)
}

package eslogger

import "time"

type Row struct {
	SchemaVersion int   `json:"schema_version"`
	MachTime      int64 `json:"mach_time"`
	EventType     int   `json:"event_type"`
	Thread        struct {
		ThreadID int `json:"thread_id"`
	} `json:"thread"`
	Version int `json:"version"`
	SeqNum  int `json:"seq_num"`
	Event   struct {
		Open struct {
			Fflag int `json:"fflag"`
			File  struct {
				Path string `json:"path"`
				Stat struct {
					StBlocks        int       `json:"st_blocks"`
					StBlksize       int       `json:"st_blksize"`
					StRdev          int       `json:"st_rdev"`
					StDev           int       `json:"st_dev"`
					StUID           int       `json:"st_uid"`
					StGid           int       `json:"st_gid"`
					StIno           int       `json:"st_ino"`
					StBirthtimespec time.Time `json:"st_birthtimespec"`
					StFlags         int       `json:"st_flags"`
					StNlink         int       `json:"st_nlink"`
					StMtimespec     time.Time `json:"st_mtimespec"`
					StCtimespec     time.Time `json:"st_ctimespec"`
					StSize          int       `json:"st_size"`
					StGen           int       `json:"st_gen"`
					StMode          int       `json:"st_mode"`
					StAtimespec     time.Time `json:"st_atimespec"`
				} `json:"stat"`
				PathTruncated bool `json:"path_truncated"`
			} `json:"file"`
		} `json:"open"`
		Exec struct {
			Script any `json:"script"`
			LastFd int `json:"last_fd"`
			Target struct {
				SigningID        string `json:"signing_id"`
				ParentAuditToken struct {
					Asid       int `json:"asid"`
					Pidversion int `json:"pidversion"`
					Ruid       int `json:"ruid"`
					Euid       int `json:"euid"`
					Rgid       int `json:"rgid"`
					Auid       int `json:"auid"`
					Egid       int `json:"egid"`
					Pid        int `json:"pid"`
				} `json:"parent_audit_token"`
				CodesigningFlags int `json:"codesigning_flags"`
				Executable       struct {
					Path string `json:"path"`
					Stat struct {
						StBlocks        int       `json:"st_blocks"`
						StBlksize       int       `json:"st_blksize"`
						StRdev          int       `json:"st_rdev"`
						StDev           int       `json:"st_dev"`
						StUID           int       `json:"st_uid"`
						StGid           int       `json:"st_gid"`
						StIno           int64     `json:"st_ino"`
						StBirthtimespec time.Time `json:"st_birthtimespec"`
						StFlags         int       `json:"st_flags"`
						StNlink         int       `json:"st_nlink"`
						StMtimespec     time.Time `json:"st_mtimespec"`
						StCtimespec     time.Time `json:"st_ctimespec"`
						StSize          int       `json:"st_size"`
						StGen           int       `json:"st_gen"`
						StMode          int       `json:"st_mode"`
						StAtimespec     time.Time `json:"st_atimespec"`
					} `json:"stat"`
					PathTruncated bool `json:"path_truncated"`
				} `json:"executable"`
				Ppid int `json:"ppid"`
				Tty  struct {
					Path string `json:"path"`
					Stat struct {
						StBlocks        int       `json:"st_blocks"`
						StBlksize       int       `json:"st_blksize"`
						StRdev          int       `json:"st_rdev"`
						StDev           int       `json:"st_dev"`
						StUID           int       `json:"st_uid"`
						StGid           int       `json:"st_gid"`
						StIno           int       `json:"st_ino"`
						StBirthtimespec time.Time `json:"st_birthtimespec"`
						StFlags         int       `json:"st_flags"`
						StNlink         int       `json:"st_nlink"`
						StMtimespec     time.Time `json:"st_mtimespec"`
						StCtimespec     time.Time `json:"st_ctimespec"`
						StSize          int       `json:"st_size"`
						StGen           int       `json:"st_gen"`
						StMode          int       `json:"st_mode"`
						StAtimespec     time.Time `json:"st_atimespec"`
					} `json:"stat"`
					PathTruncated bool `json:"path_truncated"`
				} `json:"tty"`
				StartTime        time.Time `json:"start_time"`
				IsPlatformBinary bool      `json:"is_platform_binary"`
				GroupID          int       `json:"group_id"`
				AuditToken       struct {
					Asid       int `json:"asid"`
					Pidversion int `json:"pidversion"`
					Ruid       int `json:"ruid"`
					Euid       int `json:"euid"`
					Rgid       int `json:"rgid"`
					Auid       int `json:"auid"`
					Egid       int `json:"egid"`
					Pid        int `json:"pid"`
				} `json:"audit_token"`
				IsEsClient            bool `json:"is_es_client"`
				ResponsibleAuditToken struct {
					Asid       int `json:"asid"`
					Pidversion int `json:"pidversion"`
					Ruid       int `json:"ruid"`
					Euid       int `json:"euid"`
					Rgid       int `json:"rgid"`
					Auid       int `json:"auid"`
					Egid       int `json:"egid"`
					Pid        int `json:"pid"`
				} `json:"responsible_audit_token"`
				SessionID    int    `json:"session_id"`
				OriginalPpid int    `json:"original_ppid"`
				Cdhash       string `json:"cdhash"`
				TeamID       any    `json:"team_id"`
			} `json:"target"`
			ImageCpusubtype int `json:"image_cpusubtype"`
			ImageCputype    int `json:"image_cputype"`
			Fds             []struct {
				Fd     int `json:"fd"`
				Fdtype int `json:"fdtype"`
				Pipe   struct {
					PipeID int `json:"pipe_id"`
				} `json:"pipe,omitempty"`
			} `json:"fds"`
			Env  []string `json:"env"`
			Args []string `json:"args"`
			Cwd  struct {
				Path string `json:"path"`
				Stat struct {
					StBlocks        int       `json:"st_blocks"`
					StBlksize       int       `json:"st_blksize"`
					StRdev          int       `json:"st_rdev"`
					StDev           int       `json:"st_dev"`
					StUID           int       `json:"st_uid"`
					StGid           int       `json:"st_gid"`
					StIno           int       `json:"st_ino"`
					StBirthtimespec time.Time `json:"st_birthtimespec"`
					StFlags         int       `json:"st_flags"`
					StNlink         int       `json:"st_nlink"`
					StMtimespec     time.Time `json:"st_mtimespec"`
					StCtimespec     time.Time `json:"st_ctimespec"`
					StSize          int       `json:"st_size"`
					StGen           int       `json:"st_gen"`
					StMode          int       `json:"st_mode"`
					StAtimespec     time.Time `json:"st_atimespec"`
				} `json:"stat"`
				PathTruncated bool `json:"path_truncated"`
			} `json:"cwd"`
		} `json:"exec"`
	} `json:"event"`
	Time   time.Time `json:"time"`
	Action struct {
		Result struct {
			Result struct {
				Flags int64 `json:"flags"`
			} `json:"result"`
			ResultType int `json:"result_type"`
		} `json:"result"`
	} `json:"action"`
	Process struct {
		SigningID        string `json:"signing_id"`
		ParentAuditToken struct {
			Asid       int   `json:"asid"`
			Pidversion int   `json:"pidversion"`
			Ruid       int   `json:"ruid"`
			Euid       int   `json:"euid"`
			Rgid       int   `json:"rgid"`
			Auid       int64 `json:"auid"`
			Egid       int   `json:"egid"`
			Pid        int   `json:"pid"`
		} `json:"parent_audit_token"`
		CodesigningFlags int `json:"codesigning_flags"`
		Executable       struct {
			Path string `json:"path"`
			Stat struct {
				StBlocks        int       `json:"st_blocks"`
				StBlksize       int       `json:"st_blksize"`
				StRdev          int       `json:"st_rdev"`
				StDev           int       `json:"st_dev"`
				StUID           int       `json:"st_uid"`
				StGid           int       `json:"st_gid"`
				StIno           int64     `json:"st_ino"`
				StBirthtimespec time.Time `json:"st_birthtimespec"`
				StFlags         int       `json:"st_flags"`
				StNlink         int       `json:"st_nlink"`
				StMtimespec     time.Time `json:"st_mtimespec"`
				StCtimespec     time.Time `json:"st_ctimespec"`
				StSize          int       `json:"st_size"`
				StGen           int       `json:"st_gen"`
				StMode          int       `json:"st_mode"`
				StAtimespec     time.Time `json:"st_atimespec"`
			} `json:"stat"`
			PathTruncated bool `json:"path_truncated"`
		} `json:"executable"`
		Ppid             int       `json:"ppid"`
		Tty              any       `json:"tty"`
		StartTime        time.Time `json:"start_time"`
		IsPlatformBinary bool      `json:"is_platform_binary"`
		GroupID          int       `json:"group_id"`
		AuditToken       struct {
			Asid       int `json:"asid"`
			Pidversion int `json:"pidversion"`
			Ruid       int `json:"ruid"`
			Euid       int `json:"euid"`
			Rgid       int `json:"rgid"`
			Auid       int `json:"auid"`
			Egid       int `json:"egid"`
			Pid        int `json:"pid"`
		} `json:"audit_token"`
		IsEsClient            bool `json:"is_es_client"`
		ResponsibleAuditToken struct {
			Asid       int `json:"asid"`
			Pidversion int `json:"pidversion"`
			Ruid       int `json:"ruid"`
			Euid       int `json:"euid"`
			Rgid       int `json:"rgid"`
			Auid       int `json:"auid"`
			Egid       int `json:"egid"`
			Pid        int `json:"pid"`
		} `json:"responsible_audit_token"`
		SessionID    int    `json:"session_id"`
		OriginalPpid int    `json:"original_ppid"`
		Cdhash       string `json:"cdhash"`
		TeamID       any    `json:"team_id"`
	} `json:"process"`
	ActionType   int `json:"action_type"`
	GlobalSeqNum int `json:"global_seq_num"`
}

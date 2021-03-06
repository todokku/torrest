// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-02-05 17:28:47.110396932 +0000 WET m=+0.064899008

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "i96751414",
            "url": "https://github.com/i96751414/torrest",
            "email": "i96751414@gmail.com"
        },
        "license": {
            "name": "GPL3.0",
            "url": "https://www.gnu.org/licenses/gpl-3.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/add/magnet": {
            "get": {
                "description": "add magnet to service",
                "produces": [
                    "application/json"
                ],
                "summary": "Add Magnet",
                "operationId": "add-magnet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "magnet URI",
                        "name": "uri",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.NewTorrentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/add/torrent": {
            "post": {
                "description": "add torrent file to service",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add Torrent File",
                "operationId": "add-torrent",
                "parameters": [
                    {
                        "type": "file",
                        "description": "torrent file",
                        "name": "torrent",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.NewTorrentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/pause": {
            "get": {
                "description": "pause service",
                "produces": [
                    "application/json"
                ],
                "summary": "Pause",
                "operationId": "pause",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MessageResponse"
                        }
                    }
                }
            }
        },
        "/resume": {
            "get": {
                "description": "resume service",
                "produces": [
                    "application/json"
                ],
                "summary": "Resume",
                "operationId": "resume",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MessageResponse"
                        }
                    }
                }
            }
        },
        "/settings/get": {
            "get": {
                "description": "get settings in JSON object",
                "produces": [
                    "application/json"
                ],
                "summary": "Get current settings",
                "operationId": "get-settings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/settings.Settings"
                        }
                    }
                }
            }
        },
        "/settings/set": {
            "post": {
                "description": "set settings given the provided JSON object",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Set settings",
                "operationId": "set-settings",
                "parameters": [
                    {
                        "description": "Settings to be set",
                        "name": "default",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/settings.Settings"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/settings.Settings"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/shutdown": {
            "get": {
                "description": "shutdown server",
                "summary": "Shutdown",
                "operationId": "shutdown",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "get service status",
                "produces": [
                    "application/json"
                ],
                "summary": "Status",
                "operationId": "status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/bittorrent.ServiceStatus"
                        }
                    }
                }
            }
        },
        "/torrents": {
            "get": {
                "description": "list all torrents from service",
                "produces": [
                    "application/json"
                ],
                "summary": "List Torrents",
                "operationId": "list-torrents",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "get torrents status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.TorrentInfoResponse"
                            }
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/download": {
            "get": {
                "description": "download all files from torrent",
                "produces": [
                    "application/json"
                ],
                "summary": "Download",
                "operationId": "download-torrent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MessageResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/files": {
            "get": {
                "description": "get a list of the torrent files and its details",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Torrent Files",
                "operationId": "torrent-files",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.FileInfoResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/files/{file}/download": {
            "get": {
                "description": "download file from torrent given its id",
                "produces": [
                    "application/json"
                ],
                "summary": "Download File",
                "operationId": "download-file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "file id",
                        "name": "file",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "buffer file",
                        "name": "buffer",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/files/{file}/hash": {
            "get": {
                "description": "calculate file hash suitable for opensubtitles",
                "produces": [
                    "application/json"
                ],
                "summary": "Calculate file hash",
                "operationId": "file-hash",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "file id",
                        "name": "file",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.FileHash"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/files/{file}/info": {
            "get": {
                "description": "get file info from torrent given its id",
                "produces": [
                    "application/json"
                ],
                "summary": "Get File Info",
                "operationId": "file-info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "file id",
                        "name": "file",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/bittorrent.FileInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/files/{file}/serve": {
            "get": {
                "description": "serve file from torrent given its id",
                "produces": [
                    "application/json"
                ],
                "summary": "Serve File",
                "operationId": "serve-file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "file id",
                        "name": "file",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/files/{file}/status": {
            "get": {
                "description": "get file status from torrent given its id",
                "produces": [
                    "application/json"
                ],
                "summary": "Get File Status",
                "operationId": "file-status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "file id",
                        "name": "file",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/bittorrent.FileStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/files/{file}/stop": {
            "get": {
                "description": "stop file download from torrent given its id",
                "produces": [
                    "application/json"
                ],
                "summary": "Stop File Download",
                "operationId": "stop-file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "file id",
                        "name": "file",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/pause": {
            "get": {
                "description": "pause torrent from service",
                "produces": [
                    "application/json"
                ],
                "summary": "Pause Torrent",
                "operationId": "pause-torrent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MessageResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/remove": {
            "get": {
                "description": "remove torrent from service",
                "produces": [
                    "application/json"
                ],
                "summary": "Remove Torrent",
                "operationId": "remove-torrent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "delete files",
                        "name": "delete",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MessageResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/resume": {
            "get": {
                "description": "resume a paused torrent",
                "produces": [
                    "application/json"
                ],
                "summary": "Resume Torrent",
                "operationId": "resume-torrent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MessageResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/torrents/{infoHash}/status": {
            "get": {
                "description": "get torrent status",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Torrent Status",
                "operationId": "torrent-status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "torrent info hash",
                        "name": "infoHash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/bittorrent.TorrentStatus"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Houston, we have a problem!"
                }
            }
        },
        "api.FileHash": {
            "type": "object",
            "properties": {
                "hash": {
                    "type": "string"
                }
            }
        },
        "api.FileInfoResponse": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/bittorrent.FileInfo"
            }
        },
        "api.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "done"
                }
            }
        },
        "api.NewTorrentResponse": {
            "type": "object",
            "properties": {
                "info_hash": {
                    "type": "string",
                    "example": "000102030405060708090a0b0c0d0e0f10111213"
                }
            }
        },
        "api.TorrentInfoResponse": {
            "type": "object",
            "properties": {
                "info_hash": {
                    "type": "string"
                },
                "status": {
                    "type": "object",
                    "$ref": "#/definitions/bittorrent.TorrentStatus"
                }
            }
        },
        "bittorrent.FileInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "length": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "bittorrent.FileStatus": {
            "type": "object",
            "properties": {
                "buffering_progress": {
                    "type": "number"
                },
                "progress": {
                    "type": "number"
                },
                "state": {
                    "type": "integer"
                }
            }
        },
        "bittorrent.ServiceStatus": {
            "type": "object",
            "properties": {
                "download_rate": {
                    "type": "integer"
                },
                "is_paused": {
                    "type": "boolean"
                },
                "num_torrents": {
                    "type": "integer"
                },
                "progress": {
                    "type": "number"
                },
                "upload_rate": {
                    "type": "integer"
                }
            }
        },
        "bittorrent.TorrentStatus": {
            "type": "object",
            "properties": {
                "active_time": {
                    "type": "integer"
                },
                "all_time_download": {
                    "type": "integer"
                },
                "all_time_upload": {
                    "type": "integer"
                },
                "download_rate": {
                    "type": "integer"
                },
                "finished_time": {
                    "type": "integer"
                },
                "has_metadata": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "paused": {
                    "type": "boolean"
                },
                "peers": {
                    "type": "integer"
                },
                "peers_total": {
                    "type": "integer"
                },
                "progress": {
                    "type": "number"
                },
                "seeders": {
                    "type": "integer"
                },
                "seeders_total": {
                    "type": "integer"
                },
                "seeding_time": {
                    "type": "integer"
                },
                "state": {
                    "type": "integer"
                },
                "upload_rate": {
                    "type": "integer"
                }
            }
        },
        "settings.ProxySettings": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "settings.Settings": {
            "type": "object",
            "properties": {
                "buffer_size": {
                    "type": "integer",
                    "example": 20971520
                },
                "connections_limit": {
                    "type": "integer",
                    "example": 0
                },
                "disable_dht": {
                    "type": "boolean",
                    "example": false
                },
                "disable_upnp": {
                    "type": "boolean",
                    "example": false
                },
                "download_path": {
                    "type": "string",
                    "example": "downloads"
                },
                "encryption_policy": {
                    "type": "integer",
                    "example": 0
                },
                "limit_after_buffering": {
                    "type": "boolean",
                    "example": false
                },
                "listen_interfaces": {
                    "type": "string"
                },
                "listen_port": {
                    "type": "integer",
                    "example": 6889
                },
                "max_download_rate": {
                    "type": "integer",
                    "example": 0
                },
                "max_upload_rate": {
                    "type": "integer",
                    "example": 0
                },
                "outgoing_interfaces": {
                    "type": "string"
                },
                "proxy": {
                    "type": "object",
                    "$ref": "#/definitions/settings.ProxySettings"
                },
                "seed_time_limit": {
                    "type": "integer",
                    "example": 0
                },
                "seed_time_ratio_limit": {
                    "type": "integer",
                    "example": 0
                },
                "session_save": {
                    "type": "integer",
                    "example": 30
                },
                "share_ratio_limit": {
                    "type": "integer",
                    "example": 0
                },
                "torrents_path": {
                    "type": "string",
                    "example": "downloads/Torrents"
                },
                "tuned_storage": {
                    "type": "boolean",
                    "example": false
                },
                "user_agent": {
                    "type": "integer",
                    "example": 0
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Torrest API",
	Description: "Torrent server with a REST API.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

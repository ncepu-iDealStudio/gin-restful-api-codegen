# coding: utf-8
{{ $CodeDict := . }}
from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()

class {{$CodeDict.TableInfo.StructName}}(db.Model):
    __tablename__ = '{{$CodeDict.TableInfo.TableName}}' {{range .TableInfo.Columns}}
    {{.Field}} = db.Column({{.GoType}}, primary_key={{if (eq .Key "PRI")}}True{{else}}False{{end}}, nullable={{if .Null}}True{{else}}False{{end}}, info="{{.Comment}}"){{end}}
# d-tck-app-operation-adminrole
システム運用管理

### シェル・プログラムマッピングテーブル
|No.|シェル名|プログラム名|
| ------------- |:-------------|-----|
|1|DATADEL|data-del|
|2|EC2START|ec2-start|
|3|EC2STOP|ec2-stop|
|4|EC2RCVALLSTART|start-all-instances|
|5|EC2RCVALLSTOP|stop-all-instances|
|6|LLS|send-command|
|7|LLT|send-command|
|8|SSS|start-jushin|
|9|SED|send-command|
|10|PROCESSALLSTART|start-jushin → send-command|
|11|PROCESSALLSTOP|send-command (LLT) → send-command (SED)|
|12|RCVLINE|send-command|
|13|RCVLINESTART|start-jushin|
|14|TOIAWASESTART|toiawase-start|
|15|TOIAWASESTOP|send-command|
|16|DATATEISEISTART|start-ecs|
|17|PROCESSGETDATA|process-get-data|
|18|IPPUNASHISEISEI|start-ecs|
|19|IPPUNASHITEISEI|start-ecs|
|20|IPPUNASHIRECREATE|recreate-one-min|
|21|LLX|send-command|
|22|ALLLLX|send-command|
|23|BACKUP|start-ecs|
|24|MSGRECEIVECHECK|message-receive-check|
|25|CHIKUSEKICHECK|chikuseki-check|
|26|LLD|send-command|
|27|ALLLLD|send-command|
|28|TOIAWASELLD|send-command|
|29|SSSRECOVER|作成中|
|30|TABLECREATE|create-table|
|31|DDSDATATEIKYOU|start-ecs|
|32|TKTOTAL|tktotal|
|33|KANSHIPROCESS|kanshi-process|
|34|KANSHISTOP|作成中|
|35|STATUSCHECK|show-status|
|36|UPDATESTATUS|update-status|
|37|UPDATECALENDAR|内部レビュー中|
|38|ADDREPLICA|作成中|
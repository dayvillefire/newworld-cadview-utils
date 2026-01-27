# call_report.sql
# @jbuchbinder

SET SQL_MODE='';
DROP TABLE IF EXISTS call_stats;
CREATE TABLE call_stats
SELECT um.fd_id AS 'fdid', um.station_id AS 'sta', i.incident_number, cs.call_priority AS 'pri', cs.call_type, GROUP_CONCAT(us.unit_number) AS units_arrived,
  IF(IFNULL(ss.dispatch_date_time, MIN(us.dispatch_date_time))>MIN(us.enroute_date_time),cs.dispatched_date_time,IFNULL(ss.dispatch_date_time, MIN(us.dispatch_date_time))) AS 'dispatched',
  MIN(us.enroute_date_time) AS 'enroute', MIN(us.arrive_date_time) AS 'arrived',
  TIME_TO_SEC(TIMEDIFF(STR_TO_DATE(MIN(us.enroute_date_time), '%m/%d/%Y %T'), STR_TO_DATE(IF(IFNULL(ss.dispatch_date_time, MIN(us.dispatch_date_time))>MIN(us.enroute_date_time),cs.dispatched_date_time,IFNULL(ss.dispatch_date_time, MIN(us.dispatch_date_time))), '%m/%d/%Y %T'))) AS 'enroute_time',
  TIME_TO_SEC(TIMEDIFF(STR_TO_DATE(MIN(us.arrive_date_time), '%m/%d/%Y %T'), STR_TO_DATE(IF(IFNULL(ss.dispatch_date_time, MIN(us.dispatch_date_time))>MIN(us.enroute_date_time),cs.dispatched_date_time,IFNULL(ss.dispatch_date_time, MIN(us.dispatch_date_time))), '%m/%d/%Y %T'))) AS 'arrival_time'
FROM unit_mappings um
LEFT OUTER JOIN incident_objs i ON i.ori = um.fd_id
LEFT OUTER JOIN call_objs cs ON cs.id = i.call_id
LEFT OUTER JOIN unit_objs us ON cs.id = us.call_id AND (
  us.unit_number LIKE CONCAT('%', um.station_id) AND
  NOT us.unit_number IN ( CONCAT('RES5', um.station_id), CONCAT('RES', um.station_id), CONCAT(um.station_id, 'TECH'), CONCAT(um.station_id, 'PAID'), CONCAT('STA', um.station_id), CONCAT('STA5', um.station_id) )) AND us.arrive_date_time <> ''
LEFT OUTER JOIN unit_objs ss ON cs.id = ss.call_id AND
  ss.unit_number IN ( CONCAT('STA', um.station_id), CONCAT('RES', um.station_id), CONCAT(um.station_id, 'FAST'), CONCAT(um.station_id, 'TECH'), CONCAT(um.station_id, 'PAID') )
GROUP BY um.fd_id, i.incident_number;


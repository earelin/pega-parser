CREATE VIEW concellos_datos
AS
SELECT proceso_electoral_id,
       concello_id,
       SUM(censo)        AS censo,
       SUM(votos_blanco) AS votos_blanco,
       SUM(votos_nulos)  AS votos_nulos
FROM mesas_electorais
GROUP BY proceso_electoral_id, concello_id;

CREATE VIEW concellos_votos_candidaturas
AS
SELECT proceso_electoral_id, concello_id, candidatura_id, SUM(votos) AS votos
FROM mesa_electoral_votos_candidaturas mevc
LEFT JOIN mesas_electorais me ON mevc.mesa_electoral_id = me.id
GROUP BY proceso_electoral_id, concello_id, candidatura_id;

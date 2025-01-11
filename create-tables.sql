-- DROP TABLE IF EXISTS mynews;
CREATE TABLE mynews (
    id           INT AUTO_INCREMENT NOT NULL,
    mainSite     VARCHAR(100) NOT NULL,
    articleTitle VARCHAR(500) NOT NULL,
    articleURL   VARCHAR(500) NOT NULL,
    `date`       VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`)
);

-- INSERT INTO mynews
--     (mainSite, articleTitle, articleURL, date)
-- VALUES
--     ('www.globo.com', 'São Paulo', 'https://g1.globo.com/sp/sao-paulo/', '2024-12-31'),
--     ('www.globo.com', 'Grupo faz arrastão em condomínio nos Jardins no ano novo; porteiro teria fugido juntoSão Paulog1', 'https://g1.globo.com/sp/sao-paulo/noticia/2025/01/03/grupo-faz-arrastao-em-condominio-nos-jardins-na-noite-de-ano-novo.ghtml', '2025-01-03');
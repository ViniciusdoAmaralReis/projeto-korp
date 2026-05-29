# Projeto Korp

Desafio técnico: serviço HTTP em Go com containerização, proxy reverso, monitoramento Prometheus/Grafana e automação com Ansible.

## Sobre o Projeto

Servidor HTTP escrito em **Go** que expõe:
- `GET /projeto-korp` → JSON `{"nome":"Projeto Korp","horario":"<UTC>"}`
- Métricas Prometheus em `/metrics` (contador de requisições e gauge de disponibilidade)

A aplicação roda em containers Docker orquestrados pelo **Docker Compose**:
- **Nginx** como proxy reverso (porta 80)
- **Prometheus** coleta as métricas
- **Grafana** exibe dashboard (volume de requisições e disponibilidade)

**Automação total** com **Ansible**: um único playbook instala Docker, copia arquivos, constrói imagens, sobe todos os containers e valida o serviço.

## Tecnologias

| Ferramenta | Função |
|------------|--------|
| Go 1.23 | Serviço HTTP |
| Docker / Docker Compose | Containerização e orquestração |
| Nginx | Proxy reverso |
| Prometheus | Coleta de métricas |
| Grafana | Visualização |
| Ansible | Provisionamento automatizado |

## Pré‑requisitos

- Linux (ou WSL2 no Windows) – para execução manual via Docker
- Ansible (se for usar a automação)
- Docker e Docker Compose (para execução direta)

## Execução

### 1. Com Docker Compose (manual)

```bash
git clone https://github.com/ViniciusdoAmaralReis/projeto-korp.git
cd projeto-korp
docker-compose up -d
```
Acesse:
- Serviço: http://localhost/projeto-korp
- Métricas: http://localhost:8080/metrics (direto) ou http://localhost/metrics (via Nginx)
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000 (usuário admin, senha admin)

### 2. Com Ansible (automação completa)

```bash
git clone https://github.com/ViniciusdoAmaralReis/projeto-korp.git
cd projeto-korp/ansible
ansible-playbook playbook.yml -K
```
O playbook irá:
- Instalar Docker, Docker Compose, curl e rsync
- Criar diretório /opt/projeto-korp
- Copiar todo o código e configurações
- Construir a imagem Go
- Subir os containers (Go, Nginx, Prometheus, Grafana)
- Testar o endpoint /projeto-korp e exibir a resposta JSON

## Monitoramento

Após subir o ambiente, configure o Grafana (uma única vez, manualmente):
1. Acesse http://localhost:3000
2. Adicione Data Source → Prometheus → URL http://prometheus:9090
3. Crie um dashboard com dois painéis:
  - Volume de requisições: rate(http_requests_total[1m]) (gráfico de linhas)
  - Disponibilidade: service_up (painel Stat)

## Decisões Técnicas

- Multi‑stage Dockerfile: imagem final baseada em Alpine (~15 MB), sem ferramentas de compilação.
- Nginx como proxy reverso: o servidor Go não fica exposto diretamente; o Nginx gerencia porta 80 e pode ser facilmente escalado.
- Métricas Prometheus: usando client_golang, com contador por endpoint e gauge de disponibilidade.
- Ansible idempotente: o playbook pode ser executado várias vezes sem efeitos colaterais; trata remoção de ambiente inconsistente.

## Licença
Este projeto é apenas para fins de avaliação técnica.


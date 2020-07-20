FROM gitpod/workspace-mysql

# RUN sudo apt-get update \
#  && sudo apt-get install -y \
#     go \
#  && sudo rm -rf /var/lib/apt/lists/* \
# RUN go get	"database/sql" \
#  && go get	"github.com/dailyburn/ratchet" \
#  && go get	"github.com/dailyburn/ratchet/logger" \
#  && go get	"github.com/dailyburn/ratchet/processors" \
#  && go get  "github.com/go-sql-driver/mysql" \
#  && go get	"github.com/rkulla/ratchet-examples/example1/packages"
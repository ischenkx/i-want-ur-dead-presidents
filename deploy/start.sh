bin/auth &
P1=$!
bin/billing &
P2=$!
bin/users &
P3=$!
bin/entities &
P4=$!
bin/api &
P4=$!
wait $P1 $P2 $P3 $P4
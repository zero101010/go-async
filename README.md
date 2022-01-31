# go-async

### Concorrência
- É executar várias coisas de forma simultânea mas sempre conciliando uma tarefa na outra, um exemplo é atender um telefone, deixar a chamada em espera, digitar algo e voltar para o telefone.
- Isso mostra como funciona a concorrência
### Paralelismo
- É executar as coisas de forma simultânea sem ser conciliando, um exemplo é digitar enquanto atende o telefone
- Isso mostra como é o paralelismo

### Scheduler
- Parte da thread ou concorrência responsável pelo gerenciamento de execução 

### Thead S0
- Cada thread faz um syscall para o SO, mas e cada thead consome cerca de 1mb 
### Concorrência golang
- Cara go routine consome somente 2kb
- O scheduler do go é cooperativo, isso 
#! / usr / bin / env bash
#    Use este script para testar se um determinado host / porta TCP está disponível

cmdname = $ ( basename $ 0 )

echoerr () { se [[ $ QUIET  -ne 1]] ;  então  echo  " $ @ "  1> & 2 ;  fi }

uso ()
{
    gato <<  USO > & 2
Uso:
    $ cmdname host: port [-s] [-t timeout] [- comando args]
    -h HOST | --host = Host HOST ou IP em teste
    -p PORT | --port = porta TCP PORT em teste
                                Como alternativa, você especifica o host e a porta como host: port
    -s | --strict Executa apenas o subcomando se o teste for bem-sucedido
    -q | --quiet Não envia mensagens de status
    -t TIMEOUT | --timeout = TIMEOUT
                                Tempo limite em segundos, zero sem tempo limite
    - COMMAND ARGS Executa o comando com args após o teste terminar
USO
    sair 1
}

wait_for ()
{
    if [[ $ TIMEOUT  -gt 0]] ;  então
        echoerr " $ cmdname : aguardando $ TIMEOUT segundos para $ HOST : $ PORT "
    outro
        echoerr " $ cmdname : aguardando $ HOST : $ PORT sem tempo limite "
    fi
    start_ts = $ ( data +% s )
    enquanto  :
    Faz
        if [[ $ ISBUSY  -eq 1]] ;  então
            nc -z $ HOST  $ PORT
            resultado = $?
        outro
            (echo > / dev / tcp / $ HOST / $ PORT ) > / dev / null 2> & 1
            resultado = $?
        fi
        if [[ $  -eq 0]] ;  então
            end_ts = $ ( date +% s )
            echoerr " $ cmdname : $ HOST : $ PORT está disponível após $ (( end_ts - start_ts )) segundos "
            pausa
        fi
        dormir 1
    feito
    return  $ result
}

wait_for_wrapper ()
{
    # Para suportar o SIGINT durante o tempo limite: http://unix.stackexchange.com/a/57692
    if [[ $ QUIET-  eq 1]] ;  então
        tempo limite $ BUSYTIMEFLAG  $ TIMEOUT  $ 0 --quiet --child --host = $ HOST --port = $ PORT --timeout = $ TIMEOUT  &
    outro
        tempo limite $ BUSYTIMEFLAG  $ TIMEOUT  $ 0 - filho --host = $ HOST --port = $ PORT --timeout = $ TIMEOUT  &
    fi
    PID = $!
    armadilha  " kill -INT - $ PID " INT
    espera  $ PID
    RESULTADO = $?
    if [[ $ RESULT  -ne 0]] ;  então
        echoerr " $ cmdname : timeout ocorreu depois de esperar $ TIMEOUT segundos para $ HOST : $ PORT "
    fi
    retornar  $ RESULTADO
}

# argumentos do processo
enquanto [[ $ #  -gt 0]]
Faz
    caso  " $ 1 "  em
        * : * )
        hostport = ( $ {1 //: / } )
        HOST = $ {hostport [0]}
        PORT = $ {hostport [1]}
        turno 1
        ;;
        --criança)
        CRIANÇA = 1
        turno 1
        ;;
        -q | --quieto)
        QUIET = 1
        turno 1
        ;;
        -s | --rigoroso)
        ESTRITA = 1
        turno 1
        ;;
        -h
        HOST = " $ 2 "
        if [[ $ HOST  ==  " " ]] ;  então  parta ;  fi
        turno 2
        ;;
        --host = * )
        HOST = " $ {1 # * =} "
        turno 1
        ;;
        -p)
        PORT = " $ 2 "
        if [[ $ PORT  ==  " " ]] ;  então  parta ;  fi
        turno 2
        ;;
        --port = * )
        PORT = " $ {1 # * =} "
        turno 1
        ;;
        -t)
        TIMEOUT = " $ 2 "
        if [[ $ TIMEOUT  ==  " " ]] ;  então  parta ;  fi
        turno 2
        ;;
        --timeout = * )
        TIMEOUT = " $ {1 # * =} "
        turno 1
        ;;
        -)
        mudança
        CLI = ( " $ @ " )
        pausa
        ;;
        --Socorro)
        uso
        ;;
        * )
        echoerr " Argumento desconhecido: $ 1 "
        uso
        ;;
    esac
feito

if [[ " $ HOST "  ==  " "  ||  " $ PORT "  ==  " " ]] ;  então
    echoerr " Erro: você precisa fornecer um host e uma porta para testar. "
    uso
fi

TIMEOUT = $ {TIMEOUT : - 15}
STRICT = $ {STRICT : - 0}
CHILD = $ {CHILD : - 0}
QUIET = $ {QUIET : - 0}

# verifique se o tempo limite é de busybox?
# verifique se o tempo limite é de busybox?
TIMEOUT_PATH = $ ( realpath $ ( que tempo limite ) )
if [[ $ TIMEOUT_PATH  = ~  " busybox " ]] ;  então
        ISBUSY = 1
        BUSYTIMEFLAG = " -t "
outro
        ISBUSY = 0
        BUSYTIMEFLAG = " "
fi

if [[ $ CHILD  -gt 0]] ;  então
    Esperar por
    RESULTADO = $?
    sair  $ RESULTADO
outro
    if [[ $ TIMEOUT  -gt 0]] ;  então
        wait_for_wrapper
        RESULTADO = $?
    outro
        Esperar por
        RESULTADO = $?
    fi
fi

if [[ $ CLI  ! =  " " ]] ;  então
    if [[ $ RESULT  -ne 0 &&  $ STRICT  -eq 1]] ;  então
        echoerr " $ cmdname : modo estrito, recusando-se a executar o subprocesso "
        sair  $ RESULTADO
    fi
    exec  " $ {CLI [@]} "
outro
    sair  $ RESULTADO
fi
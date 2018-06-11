# bash completion for dt                                   -*- shell-script -*-

__debug()
{
    if [[ -n ${BASH_COMP_DEBUG_FILE} ]]; then
        echo "$*" >> "${BASH_COMP_DEBUG_FILE}"
    fi
}

# Homebrew on Macs have version 1.3 of bash-completion which doesn't include
# _init_completion. This is a very minimal version of that function.
__my_init_completion()
{
    COMPREPLY=()
    _get_comp_words_by_ref "$@" cur prev words cword
}

__index_of_word()
{
    local w word=$1
    shift
    index=0
    for w in "$@"; do
        [[ $w = "$word" ]] && return
        index=$((index+1))
    done
    index=-1
}

__contains_word()
{
    local w word=$1; shift
    for w in "$@"; do
        [[ $w = "$word" ]] && return
    done
    return 1
}

__handle_reply()
{
    __debug "${FUNCNAME[0]}"
    case $cur in
        -*)
            if [[ $(type -t compopt) = "builtin" ]]; then
                compopt -o nospace
            fi
            local allflags
            if [ ${#must_have_one_flag[@]} -ne 0 ]; then
                allflags=("${must_have_one_flag[@]}")
            else
                allflags=("${flags[*]} ${two_word_flags[*]}")
            fi
            COMPREPLY=( $(compgen -W "${allflags[*]}" -- "$cur") )
            if [[ $(type -t compopt) = "builtin" ]]; then
                [[ "${COMPREPLY[0]}" == *= ]] || compopt +o nospace
            fi

            # complete after --flag=abc
            if [[ $cur == *=* ]]; then
                if [[ $(type -t compopt) = "builtin" ]]; then
                    compopt +o nospace
                fi

                local index flag
                flag="${cur%%=*}"
                __index_of_word "${flag}" "${flags_with_completion[@]}"
                COMPREPLY=()
                if [[ ${index} -ge 0 ]]; then
                    PREFIX=""
                    cur="${cur#*=}"
                    ${flags_completion[${index}]}
                    if [ -n "${ZSH_VERSION}" ]; then
                        # zsh completion needs --flag= prefix
                        eval "COMPREPLY=( \"\${COMPREPLY[@]/#/${flag}=}\" )"
                    fi
                fi
            fi
            return 0;
            ;;
    esac

    # check if we are handling a flag with special work handling
    local index
    __index_of_word "${prev}" "${flags_with_completion[@]}"
    if [[ ${index} -ge 0 ]]; then
        ${flags_completion[${index}]}
        return
    fi

    # we are parsing a flag and don't have a special handler, no completion
    if [[ ${cur} != "${words[cword]}" ]]; then
        return
    fi

    local completions
    completions=("${commands[@]}")
    if [[ ${#must_have_one_noun[@]} -ne 0 ]]; then
        completions=("${must_have_one_noun[@]}")
    fi
    if [[ ${#must_have_one_flag[@]} -ne 0 ]]; then
        completions+=("${must_have_one_flag[@]}")
    fi
    COMPREPLY=( $(compgen -W "${completions[*]}" -- "$cur") )

    if [[ ${#COMPREPLY[@]} -eq 0 && ${#noun_aliases[@]} -gt 0 && ${#must_have_one_noun[@]} -ne 0 ]]; then
        COMPREPLY=( $(compgen -W "${noun_aliases[*]}" -- "$cur") )
    fi

    if [[ ${#COMPREPLY[@]} -eq 0 ]]; then
        declare -F __custom_func >/dev/null && __custom_func
    fi

    # available in bash-completion >= 2, not always present on macOS
    if declare -F __ltrim_colon_completions >/dev/null; then
        __ltrim_colon_completions "$cur"
    fi
}

# The arguments should be in the form "ext1|ext2|extn"
__handle_filename_extension_flag()
{
    local ext="$1"
    _filedir "@(${ext})"
}

__handle_subdirs_in_dir_flag()
{
    local dir="$1"
    pushd "${dir}" >/dev/null 2>&1 && _filedir -d && popd >/dev/null 2>&1
}

__handle_flag()
{
    __debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    # if a command required a flag, and we found it, unset must_have_one_flag()
    local flagname=${words[c]}
    local flagvalue
    # if the word contained an =
    if [[ ${words[c]} == *"="* ]]; then
        flagvalue=${flagname#*=} # take in as flagvalue after the =
        flagname=${flagname%%=*} # strip everything after the =
        flagname="${flagname}=" # but put the = back
    fi
    __debug "${FUNCNAME[0]}: looking for ${flagname}"
    if __contains_word "${flagname}" "${must_have_one_flag[@]}"; then
        must_have_one_flag=()
    fi

    # if you set a flag which only applies to this command, don't show subcommands
    if __contains_word "${flagname}" "${local_nonpersistent_flags[@]}"; then
      commands=()
    fi

    # keep flag value with flagname as flaghash
    if [ -n "${flagvalue}" ] ; then
        flaghash[${flagname}]=${flagvalue}
    elif [ -n "${words[ $((c+1)) ]}" ] ; then
        flaghash[${flagname}]=${words[ $((c+1)) ]}
    else
        flaghash[${flagname}]="true" # pad "true" for bool flag
    fi

    # skip the argument to a two word flag
    if __contains_word "${words[c]}" "${two_word_flags[@]}"; then
        c=$((c+1))
        # if we are looking for a flags value, don't show commands
        if [[ $c -eq $cword ]]; then
            commands=()
        fi
    fi

    c=$((c+1))

}

__handle_noun()
{
    __debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    if __contains_word "${words[c]}" "${must_have_one_noun[@]}"; then
        must_have_one_noun=()
    elif __contains_word "${words[c]}" "${noun_aliases[@]}"; then
        must_have_one_noun=()
    fi

    nouns+=("${words[c]}")
    c=$((c+1))
}

__handle_command()
{
    __debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    local next_command
    if [[ -n ${last_command} ]]; then
        next_command="_${last_command}_${words[c]//:/__}"
    else
        if [[ $c -eq 0 ]]; then
            next_command="_$(basename "${words[c]//:/__}")"
        else
            next_command="_${words[c]//:/__}"
        fi
    fi
    c=$((c+1))
    __debug "${FUNCNAME[0]}: looking for ${next_command}"
    declare -F "$next_command" >/dev/null && $next_command
}

__handle_word()
{
    if [[ $c -ge $cword ]]; then
        __handle_reply
        return
    fi
    __debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"
    if [[ "${words[c]}" == -* ]]; then
        __handle_flag
    elif __contains_word "${words[c]}" "${commands[@]}"; then
        __handle_command
    elif [[ $c -eq 0 ]] && __contains_word "$(basename "${words[c]}")" "${commands[@]}"; then
        __handle_command
    else
        __handle_noun
    fi
    __handle_word
}

_dt_base64_decode()
{
    last_command="dt_base64_decode"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    flags+=("--format=")
    two_word_flags+=("-f")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_base64_encode()
{
    last_command="dt_base64_encode"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    flags+=("--format=")
    two_word_flags+=("-f")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_base64()
{
    last_command="dt_base64"
    commands=()
    commands+=("decode")
    commands+=("encode")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--format=")
    two_word_flags+=("-f")
    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_completion()
{
    last_command="dt_completion"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--help")
    flags+=("-h")
    local_nonpersistent_flags+=("--help")
    flags+=("--shell=")
    two_word_flags+=("-s")
    local_nonpersistent_flags+=("--shell=")
    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_date()
{
    last_command="dt_date"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_hash()
{
    last_command="dt_hash"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--algo=")
    two_word_flags+=("-a")
    local_nonpersistent_flags+=("--algo=")
    flags+=("--cost=")
    two_word_flags+=("-c")
    local_nonpersistent_flags+=("--cost=")
    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_html()
{
    last_command="dt_html"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--escape=")
    two_word_flags+=("-e")
    local_nonpersistent_flags+=("--escape=")
    flags+=("--escapeFile=")
    two_word_flags+=("-f")
    local_nonpersistent_flags+=("--escapeFile=")
    flags+=("--unescape=")
    two_word_flags+=("-d")
    local_nonpersistent_flags+=("--unescape=")
    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_jwt()
{
    last_command="dt_jwt"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--base64Secret")
    flags+=("-b")
    local_nonpersistent_flags+=("--base64Secret")
    flags+=("--secret=")
    two_word_flags+=("-s")
    local_nonpersistent_flags+=("--secret=")
    flags+=("--secretFile=")
    two_word_flags+=("-f")
    local_nonpersistent_flags+=("--secretFile=")
    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_random()
{
    last_command="dt_random"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--algo=")
    two_word_flags+=("-a")
    local_nonpersistent_flags+=("--algo=")
    flags+=("--length=")
    two_word_flags+=("-l")
    local_nonpersistent_flags+=("--length=")
    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_server()
{
    last_command="dt_server"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--address=")
    two_word_flags+=("-a")
    local_nonpersistent_flags+=("--address=")
    flags+=("--open")
    flags+=("-o")
    local_nonpersistent_flags+=("--open")
    flags+=("--port=")
    two_word_flags+=("-p")
    local_nonpersistent_flags+=("--port=")
    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_uri_decode()
{
    last_command="dt_uri_decode"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_uri_encode()
{
    last_command="dt_uri_encode"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_uri()
{
    last_command="dt_uri"
    commands=()
    commands+=("decode")
    commands+=("encode")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_uuid()
{
    last_command="dt_uuid"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--domain=")
    two_word_flags+=("-d")
    local_nonpersistent_flags+=("--domain=")
    flags+=("--namespace=")
    two_word_flags+=("-n")
    local_nonpersistent_flags+=("--namespace=")
    flags+=("--version=")
    two_word_flags+=("-v")
    local_nonpersistent_flags+=("--version=")
    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt_version()
{
    last_command="dt_version"
    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dt()
{
    last_command="dt"
    commands=()
    commands+=("base64")
    commands+=("completion")
    commands+=("date")
    commands+=("hash")
    commands+=("html")
    commands+=("jwt")
    commands+=("random")
    commands+=("server")
    commands+=("uri")
    commands+=("uuid")
    commands+=("version")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

__start_dt()
{
    local cur prev words cword
    declare -A flaghash 2>/dev/null || :
    if declare -F _init_completion >/dev/null 2>&1; then
        _init_completion -s || return
    else
        __my_init_completion -n "=" || return
    fi

    local c=0
    local flags=()
    local two_word_flags=()
    local local_nonpersistent_flags=()
    local flags_with_completion=()
    local flags_completion=()
    local commands=("dt")
    local must_have_one_flag=()
    local must_have_one_noun=()
    local last_command
    local nouns=()

    __handle_word
}

if [[ $(type -t compopt) = "builtin" ]]; then
    complete -o default -F __start_dt dt
else
    complete -o default -o nospace -F __start_dt dt
fi

# ex: ts=4 sw=4 et filetype=sh

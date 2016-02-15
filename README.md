alfred-trello-toggl
==========

This workflow assist you to be lazy for your tasks management :)
the flow of tasking is like Kanban structure.
You just type few words, start toggl timer and move your cards.

# Install

You can download in the releases page and double click it

# Usage

## setting

At the first time, you have to set your toggl and trello account.

- toggl
    - login
    - token
- trello
    - apikey
    - apisecret

next, create lists in your board like this

- new
    - when you create task
- progress
    - when you start the task
- wait
    - when your task become other's
- done
    - when your task done


## start

### Show all tasks

- [command] tasks
    - search task from trello
    - show all tasks
- [command] tasks [task]
    - create task if it's undefined
    - status to "new"

## edit

- [comamnd] tasks [task]
    - start timer if stopped
    - status to "progress"
- [comamnd] tasks [task]
    - stop timer if started
    - stauts to "done"

## sync

if you put tasks another way, you have to reset the task cache.

- [command] sync
    
## comming soon

### waiting

When your task becomes other person's, you can postpone the task

- [comamnd] wait [task]
    - stop timer if started
    - status to "wait"

### change board name


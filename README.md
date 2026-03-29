# ModalsDB - Graph module

ModalsDB graph module has basic functions to call itself graph database.

# Queries

This modules gives basic set of queries which can be made

## Select
- Select all incoming relations to node: <code>SELECT RELATIONS IN [john, alice];</code>
- Select all outgoing relations from node: <code>SELECT RELATIONS OUT [john, alice];</code>
- Select all relations (outgoind and incoming): <code>SELECT RELATIONS ALL [john, alice];</code>

## Insert
- Insert new node: <code>INSERT john;</code>
- Insert relation going in one direction: <code>INSERT john->[:works_for]->google;</code>
- Insert relation that goes in both directions: <code>INSERT pizza<->[:has,:on_top]<->cheese;</code>

## Count
- Count how many relations of given type goes into node: <code>COUNT IN john;</code>

## Check
- Check if there is relations between nodes: <code>CHECK john AND alice;</code>
- Check if there is a path between nodes: <code>CHECK PATH BETWEEN john AND alice;</code>
- Check if there is a certain relation between nodes: <code>CHECK john AND alice HAS [:works_on,:loves]</code>

## Update

- Update data in the node: <code>UPDATE NODE DATA john WITH {'key1': 1, 'key2': 'hello world!'}</code>
- Update relation data: <code>UPDATE RELATION DATA BETWEEN john AND alice WITH {'key1': 1, 'key2': 'hello world!'}</code>

## Delete

- Delete relation between two nodes: <code>DELETE RELATION BETWEEN john AND alice;</code>
- Delete node from database: <code>DELETE john;</code>
- In case of wrongful deletion, for BACKOFF_TIMEOUT seconds, there is: <code>RESTORE john;</code>
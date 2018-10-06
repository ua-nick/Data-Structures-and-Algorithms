/*
heap.c
July 8th, 2018
Braelyn Rotman
*/
#include "heap.h"
#include "heapADT.h"

#include <string.h>
#include <stdio.h>
#include <stdbool.h>

void printArray (Heap *heap);
void destroyData(void *data);
void printNode(void *toBePrinted);
int compare(const void *first, const void *second);

Heap *createHeap(size_t initialSize, HEAP_TYPE htype, void (*destroyData)(void *data),void (*printNode)(void *toBePrinted),int (*compare)(const void *first, const void *second)){
  Heap *newHeap = malloc(sizeof(Heap));
  newHeap->type = htype;
  /*newHeap->size = initialSize;*/
  newHeap->destroyData = destroyData;
  newHeap->printNode = printNode;
  newHeap->compare = compare;
  newHeap->arr = calloc(initialSize, sizeof(HeapNode *));
  newHeap->lastPosition = NULL;
  newHeap->size = initialSize;

  return newHeap;
}

HeapNode *createHeapNode(void *data){
  HeapNode *newNode = malloc(sizeof(HeapNode));
  newNode->left = NULL; /* child node */
  newNode->right = NULL; /* child node */
  newNode->parent = NULL;
  newNode->data = data;

  return newNode;
}

void insertHeapNode(Heap *heap, void *data){
  HeapNode *insert = createHeapNode(data);
  int i = 0;
  if (heap->lastPosition == NULL) /*insert root */
  {
    heap->arr[0] = insert;
    heap->lastPosition = insert;
  }
  else
  {
    while (heap->arr[i] != NULL) /*find last position */
    {
      if (i+1 >= heap->size)
      {
      /*  out of space*/

        heap->arr = realloc(heap->arr, sizeof(HeapNode *)*(heap->size+2));

        for (int d = 0 ; d < 2 ; d++)
        {
          heap->arr[heap->size+d] = NULL;
        }

        heap->size = heap->size + 2;

      }
      else
      {
        i++;
      }
    }
      /* copy newNode into array */
    heap->arr[i] = insert;
    heap->lastPosition = insert;
    heapify(heap, insert);
  }


}

void deleteMinOrMax(Heap *heap){
  HeapNode *temp = heap->arr[0]; /* root node */
  HeapNode *outOfPlacePtr;
  HeapNode *tempNode;
  bool left = false, right = false, checkL = false;
  int current;
  int i = 0;

  heap->arr[0] = heap->lastPosition;
  outOfPlacePtr = heap->arr[0];

  /*remove node in last position */
  heap->destroyData(temp);

  /*********************************************************/
  /* delete last node */
  while (heap->arr[i] != NULL)
  {
    i++;
  }
  i--;
  heap->arr[i] = NULL;
  i--;
  heap->lastPosition = heap->arr[i];

  current = 0;

  if (heap->type == 1)
  {
    if (heap->arr[2*current + 1] != NULL)
    {
      checkL = true;
      if(heap->compare(outOfPlacePtr->data, heap->arr[2*current + 1]->data) == -1)
      {
        left = true;
      }
    }
    if (heap->arr[2*current + 2] != NULL)
    {
      if(heap->compare(outOfPlacePtr->data, heap->arr[2*current + 2]->data) == -1)
      {
        right = true;
      }
    }

    while (checkL == true)
    {
      if (left ==true && right == true)
      {
        if (heap->arr[2*current + 1] > heap->arr[2*current + 2])
        {

          heap->arr[current] = heap->arr[2*current + 2];
          heap->arr[2*current +2] = outOfPlacePtr;

          current = 2*current +2;
        }
        else
        {
          heap->arr[current] = heap->arr[2*current + 1];
          heap->arr[2*current +1] = outOfPlacePtr;

          current = 2*current +1;
        }
      }
      else if(left == true && right == false)
      {
        heap->arr[current] = heap->arr[2*current + 1];
        heap->arr[2*current +1] = outOfPlacePtr;

        current = 2*current +1;
      }

      checkL =false;
      left = false;
      right =false;
      if (heap->arr[2*current + 1] != NULL)
      {
        checkL = true;
        if(heap->compare(outOfPlacePtr->data, heap->arr[2*current + 1]->data) == -1)
        {
          left = true;
        }
      }
      if (heap->arr[2*current + 2] != NULL)
      {
        if(heap->compare(outOfPlacePtr->data, heap->arr[2*current + 2]->data) == -1)
        {
          right = true;
        }
      }
    }
  }
  else
  {
    if (heap->arr[2*current + 1] != NULL)
    {
      checkL = true;
      if(heap->compare(outOfPlacePtr->data, heap->arr[2*current + 1]->data) == 1)
      {
        left = true;
      }
    }
    if (heap->arr[2*current + 2] != NULL)
    {
      if(heap->compare(outOfPlacePtr->data, heap->arr[2*current + 2]->data) == 1)
      {
        right = true;
      }
    }

    while (checkL == true)
    {
      if (left ==true && right == true)
      {
        if (heap->arr[2*current + 1] > heap->arr[2*current + 2])
        {

          heap->arr[current] = heap->arr[2*current + 2];
          heap->arr[2*current +2] = outOfPlacePtr;

          current = 2*current +2;
        }
        else
        {
          heap->arr[current] = heap->arr[2*current + 1];
          heap->arr[2*current +1] = outOfPlacePtr;

          current = 2*current +1;
        }
      }
      else if(left == true && right == false)
      {
        heap->arr[current] = heap->arr[2*current + 1];
        heap->arr[2*current +1] = outOfPlacePtr;

        current = 2*current +1;
      }

      checkL =false;
      left = false;
      right =false;
      if (heap->arr[2*current + 1] != NULL)
      {
        checkL = true;
        if(heap->compare(outOfPlacePtr->data, heap->arr[2*current + 1]->data) == 1)
        {
          left = true;
        }
      }
      if (heap->arr[2*current + 2] != NULL)
      {
        if(heap->compare(outOfPlacePtr->data, heap->arr[2*current + 2]->data) == 1)
        {
          right = true;
        }
      }
    }
  }

  tempNode = heap->arr[0];
  while (tempNode != NULL)
  {
    i++;
    tempNode = heap->arr[i];
  }
  heap->lastPosition = heap->arr[i];

  /**********************************************************/
  /*  return temp->data;*/
}

void *getMinOrMax(Heap *heap){

  return heap->arr[0];
}

void changeHeapType(Heap *heap){
    HeapNode *current = heap->arr[0];
    Heap *localHeap;
    int i = 0;
    int d =0;

    if (heap->type == 0)
    {
      heap->type = 1;
    }
    else
    {
      heap->type = 0;
    }
    localHeap = createHeap(10, heap->type, &destroyData, &printNode, &compare);

    while (current != NULL)
    {
      insertHeapNode(localHeap, current->data);
      i++;
      current = heap->arr[i];
    }

    d = 0;
    while (heap->arr[d] != NULL)
    {
      free(heap->arr[d]);
      heap->arr[d] = localHeap->arr[d];
      d++;
    }
    free(localHeap);
}

void deleteHeap(Heap *heap){
  int i = 0;
  while (heap->arr[i] != NULL)
  {
    heap->destroyData((void *)heap->arr[i]);
    free(heap->arr[i]);
    i++;
  }

  free(heap);

}

void heapify (Heap *heap, HeapNode * newNode){
  int parentNodeindex; /* = newNode->parent;*/
  HeapNode *temp;
  int i = 0;
  int pos = 0;

  while (heap->compare(heap->arr[pos]->data, newNode->data) != 0)
  {
    pos++;
  }

  parentNodeindex = (pos - 1) /2;

  if (heap->type == 1)
  {
    while (heap->compare(newNode->data, heap->arr[parentNodeindex]->data) == 1) /*max heap */
    {
      temp = heap->arr[parentNodeindex];
      heap->arr[parentNodeindex] = newNode;
      heap->arr[pos] = temp;

      pos = (pos - 1) /2;
      parentNodeindex = (pos - 1)/2;
    }
  }
  else
  {
    while (heap->compare(newNode->data, heap->arr[parentNodeindex]->data) == -1) /*min heap */
    {
      temp = heap->arr[parentNodeindex];
      heap->arr[parentNodeindex] = newNode;
      heap->arr[pos] = temp;
    }
  }
  temp = heap->arr[0];
  while (temp != NULL)
  {
    i++;
    temp = heap->arr[i];
  }
  i--;
  heap->lastPosition = heap->arr[i];
}

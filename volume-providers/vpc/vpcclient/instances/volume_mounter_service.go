/*******************************************************************************
 * IBM Confidential
 * OCO Source Materials
 * IBM Cloud Container Service, 5737-D43
 * (C) Copyright IBM Corp. 2018, 2019 All Rights Reserved.
 * The source code for this program is not  published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with
 * the U.S. Copyright Office.
 ******************************************************************************/

package instances

import (
	"github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/vpc/vpcclient/client"
	"github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/vpc/vpcclient/models"
)

// VolumeMountManager operations
type VolumeMountManager interface {
	// Create the volume with authorisation by passing required information in the volume object
	AttachVolume(*models.VolumeAttachment) (*models.VolumeAttachment, error)

	// Delete the volume
	DetachVolume(volumeAttachmentID string) error
}

// VolumeMountService ...
type VolumeMountService struct {
	client client.SessionClient
}

var _ VolumeMountManager = &VolumeMountService{}

// New ...
func New(client client.SessionClient) VolumeMountManager {
	return &VolumeMountService{
		client: client,
	}
}